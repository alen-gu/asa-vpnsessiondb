package sqlite

import (
	"testing"
)

//func Test_initdb(t *testing.T) {
//	err := init_database()
//	t.Log(err)
//}

//func Test_insert_database(t *testing.T) {
//	VpnSessionDB := make(map[int]vpnsession)
//	VpnSessionDB[0] = vpnsession{crasIndex: 1101,
//		crasUsername:        "20150073",
//		crasGroup:           "vpn",
//		crasSessionDuration: 1231,
//		crasLocalAddress:    "202.120.95.1",
//		crasISPAddress:      "2.123.123.123"}
//	VpnSessionDB[1] = vpnsession{crasIndex: 1102,
//		crasUsername:        "20120312",
//		crasGroup:           "vpn",
//		crasSessionDuration: 1231,
//		crasLocalAddress:    "202.120.95.15",
//		crasISPAddress:      "1.123.123.123"}
//	VpnSessionDB[2] = vpnsession{crasIndex: 1103,
//		crasUsername:        "20150073",
//		crasGroup:           "vpn",
//		crasSessionDuration: 1231,
//		crasLocalAddress:    "202.120.95.103",
//		crasISPAddress:      "222.123.123.123"}
//	err := Insert_DB(VpnSessionDB)
//	t.Log(err)
//}

//func Test_deleteold_database(t *testing.T) {
//	err := DeleteOld_DB(60)
//	t.Log(err)
//}

func Test_search_database(t *testing.T) {
	//	sessions, err := Search_DB("20150073", 0)
	//	t.Log(sessions)
	//	t.Log(err)
	sessions, err := Search_DB("1101", 3)
	t.Log(sessions)
	t.Log(err)
	aaa, err1 := SearchByIndex_DB(1101)
	t.Log(err1)
	if aaa {
		err2 := UpdateTime_DB(1101)
		t.Log(err2)
	}
	sessions, err = Search_DB("1101", 3)
	t.Log(sessions)
	t.Log(err)
	//	sessions, err = Search_DB("202.120.95.15", 1)
	//	t.Log(sessions)
	//	t.Log(err)
	//	sessions, err = Search_DB("222.123.123.123", 2)
	//	t.Log(sessions)
	//	t.Log(err)
	//sessions, err := SearchAll_DB()
	//t.Log(sessions)
	//t.Log(err)
}
