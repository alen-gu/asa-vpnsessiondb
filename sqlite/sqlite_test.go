package sqlite

import (
	"testing"
)

/*
func Test_initdb(t *testing.T) {
	err := init_database()
	t.Log(err)
}
*/
func Test_insert_database(t *testing.T) {
	VpnSessionDB := make(map[int]vpnsession)
	VpnSessionDB[0] = vpnsession{crasUsername: "20150073",
		crasGroup:           "vpn",
		crasSessionDuration: 1231,
		crasLocalAddress:    "202.120.95.1",
		crasISPAddress:      "2.123.123.123"}
	VpnSessionDB[1] = vpnsession{crasUsername: "20120312",
		crasGroup:           "vpn",
		crasSessionDuration: 1231,
		crasLocalAddress:    "202.120.95.15",
		crasISPAddress:      "1.123.123.123"}
	VpnSessionDB[2] = vpnsession{crasUsername: "20150073",
		crasGroup:           "vpn",
		crasSessionDuration: 1231,
		crasLocalAddress:    "202.120.95.103",
		crasISPAddress:      "222.123.123.123"}
	err := insert_database(VpnSessionDB)
	t.Log(err)
}

func Test_deleteold_database(t *testing.T) {
	err := deleteold_database(60)
	t.Log(err)
}

func Test_search_database(t *testing.T) {
	sessions, err := search_database("20150073", 0)
	t.Log(sessions)
	t.Log(err)
	sessions, err = search_database("202.120.95.15", 1)
	t.Log(sessions)
	t.Log(err)
	sessions, err = search_database("222.123.123.123", 2)
	t.Log(sessions)
	t.Log(err)
	sessions, err = searchall_database()
	t.Log(sessions)
	t.Log(err)
}
