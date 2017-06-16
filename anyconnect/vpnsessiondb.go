package anyconnect

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gaochao1/gosnmp"
)

var crasSessionEntry = map[string]string{
	"crasGroup":                 "1.3.6.1.4.1.9.9.392.1.3.21.1.2",
	"crasSessionDuration":       "1.3.6.1.4.1.9.9.392.1.3.21.1.6",
	"crasLocalAddress":          "1.3.6.1.4.1.9.9.392.1.3.21.1.8",
	"crasISPAddress":            "1.3.6.1.4.1.9.9.392.1.3.21.1.10",
	"crasClientVendorString":    "1.3.6.1.4.1.9.9.392.1.3.21.1.17",
	"crasClientVersionString":   "1.3.6.1.4.1.9.9.392.1.3.21.1.18",
	"crasClientOSVendorString":  "1.3.6.1.4.1.9.9.392.1.3.21.1.19",
	"crasClientOSVersionString": "1.3.6.1.4.1.9.9.392.1.3.21.1.20",
	"crasSessionInPkts":         "1.3.6.1.4.1.9.9.392.1.3.21.1.31",
	"crasSessionOutPkts":        "1.3.6.1.4.1.9.9.392.1.3.21.1.32",
	"crasSessionInDropPkts":     "1.3.6.1.4.1.9.9.392.1.3.21.1.33",
	"crasSessionOutDropPkts":    "1.3.6.1.4.1.9.9.392.1.3.21.1.34",
	"crasSessionInOctets":       "1.3.6.1.4.1.9.9.392.1.3.21.1.35",
	"crasSessionOutOctets":      "1.3.6.1.4.1.9.9.392.1.3.21.1.36",
	"crasSessionState":          "1.3.6.1.4.1.9.9.392.1.3.21.1.37",
}

type vpnsession struct {
	crasIndex                 int64 //增加这个
	crasUsername              string
	crasGroup                 string
	crasSessionDuration       int64
	crasLocalAddress          string
	crasISPAddress            string
	crasClientVendorString    string
	crasClientVersionString   string
	crasClientOSVendorString  string
	crasClientOSVersionString string
	crasSessionInPkts         int64
	crasSessionOutPkts        int64
	crasSessionInDropPkts     int64
	crasSessionOutDropPkts    int64
	crasSessionInOctets       int64
	crasSessionOutOctets      int64
	crasSessionState          int64
}

func GetVpnSessionDB(ip, community string, crasSet []string, timeout int, retry int, limitConn int) (VpnSessionDB map[int]vpnsession) {
	var limitCh chan bool
	VpnSessionDB = make(map[int]vpnsession)
	var session vpnsession
	if limitConn > 0 {
		limitCh = make(chan bool, limitConn)
	} else {
		limitCh = make(chan bool, 1)
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println(ip+" Recovered in GetVpnSessionDB", r)
		}
	}()
	crasDefault := []string{"crasGroup"}
	crasList := append(crasDefault, crasSet...)

	chs := make([]chan map[string]interface{}, len(crasList))
	crasEntryMapList := make([]map[string]interface{}, len(crasList))

	for i, MibObject := range crasList {
		limitCh <- true
		chs[i] = make(chan map[string]interface{})
		go GetcrasSessionEntry(ip, MibObject, community, timeout, retry, chs[i], limitCh)
		time.Sleep(5 * time.Millisecond)
	}
	for i, ch := range chs {
		crasEntryMapList[i] = <-ch
	}
	if len(crasEntryMapList[0]) == 0 {
		return
	}
	for name, value := range crasEntryMapList[0] {
		index, crasUsername := CalcUsernameIndex(name)
		session.crasIndex = int64(index) //2017-6-14
		session.crasUsername = crasUsername
		session.crasGroup = value.(string)
		for i, crasEntryMap := range crasEntryMapList {
			if len(crasEntryMap) == 0 {
				continue
			}
			if i == 0 {
				continue
			}
			for _, value := range crasEntryMap {
				switch crasList[i] {
				case "crasSessionDuration":
					session.crasSessionDuration = int64(value.(int))
				case "crasLocalAddress":
					session.crasLocalAddress = value.(string)
				case "crasISPAddress":
					session.crasISPAddress = value.(string)
				case "crasClientVendorString":
					session.crasClientVendorString = value.(string)
				case "crasClientVersionString":
					session.crasClientVersionString = value.(string)
				case "crasClientOSVendorString":
					session.crasClientOSVendorString = value.(string)
				case "crasClientOSVersionString":
					session.crasClientOSVersionString = value.(string)
				case "crasSessionInPkts":
					session.crasSessionInPkts = int64(value.(uint64))
				case "crasSessionOutPkts":
					session.crasSessionOutPkts = int64(value.(uint64))
				case "crasSessionInDropPkts":
					session.crasSessionInDropPkts = int64(value.(uint64))
				case "crasSessionOutDropPkts":
					session.crasSessionOutDropPkts = int64(value.(uint64))
				case "crasSessionInOctets":
					session.crasSessionInOctets = int64(value.(uint64))
				case "crasSessionOutOctets":
					session.crasSessionOutOctets = int64(value.(uint64))
				case "crasSessionState":
					session.crasSessionState = int64(value.(int))
				}
			}
		}
		VpnSessionDB[index] = session
	}

	return
}

func CalcUsernameIndex(crasUserASCII string) (index int, crasUsername string) {
	crasUserArray := strings.Split(crasUserASCII, ".")
	crasUsername = ""
	for i, ascii := range crasUserArray {
		if i == 0 {
			continue
		}
		if i == (len(crasUserArray) - 1) {
			index, _ = strconv.Atoi(ascii)
			continue
		}
		ascii_int, _ := strconv.Atoi(ascii)
		char := rune(int(ascii_int))
		crasUsername = crasUsername + string(char)
	}
	return
}

func GetcrasSessionEntry(ip, MibObject, community string, timeout int, retry int, ch chan map[string]interface{}, limitCh chan bool) {
	var snmpPDUs []gosnmp.SnmpPDU
	oid := crasSessionEntry[MibObject]
	oidPrefix := "." + oid + "."
	crasEntry := make(map[string]interface{})

	snmpPDUs, err := RunSnmpwalk(ip, community, oid, retry, timeout)
	if err != nil {
		close(ch)
		<-limitCh
		log.Println(ip, MibObject, oid, err)
		return
	}
	for _, snmpPDU := range snmpPDUs {
		crasUserASCII := strings.Replace(snmpPDU.Name, oidPrefix, "", 1)
		crasEntry[crasUserASCII] = snmpPDU.Value
	}
	<-limitCh
	ch <- crasEntry
	return
}
