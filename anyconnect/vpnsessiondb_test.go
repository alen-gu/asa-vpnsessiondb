package anyconnect

import (
	"testing"
)

const (
	ip        = "1.1.1.1"
	community = "ecnu-test"
)

func Test_runsnmp(t *testing.T) {
	crasSet := []string{
		"caonima",
		"crasSessionDuration",
		"crasLocalAddress",
		"crasISPAddress",
		"crasClientVendorString",
		"crasClientVersionString",
		"crasClientOSVendorString",
		"crasClientOSVersionString",
		"crasSessionInPkts",
		"crasSessionOutPkts",
		"crasSessionInDropPkts",
		"crasSessionOutDropPkts",
		"crasSessionInOctets",
		"crasSessionOutOctets",
		"crasSessionState",
	}
	r1 := GetVpnSessionDB(ip, community, crasSet, 3, 3, 1)
	t.Log(r1)
}
