package cron

import (
	"log"
	"math/rand"
	//"time"
	"github.com/alen-gu/asa-vpnsessiondb/sqlite"
	"github.com/robfig/cron"
)

func UpdateDB() {
	c := cron.New()
	//spec := "*/5 * * * * ?"
	//	c.AddFunc("@every 2m", func() {
	//		VpnSessionDB := make(map[int]sqlite.VpnSession)
	//		VpnSessionDB[0] = sqlite.VpnSession{CrasIndex: int64(2000 + rand.Intn(100)),
	//			CrasUsername:        "test111111",
	//			CrasGroup:           "vpn",
	//			CrasSessionDuration: 1231,
	//			CrasLocalAddress:    "202.120.95.1",
	//			CrasISPAddress:      "2.123.123.123"}
	//		err := sqlite.Insert_DB(VpnSessionDB)
	//		//log.Println("2m", err)
	//		sessions, _ := sqlite.Search_DB("1104", 3)
	//		log.Println("第一个函数", sessions, err)

	//	})
	c.AddFunc("@every 1m", func() {
		VpnSessionDB := make(map[int]sqlite.VpnSession)
		VpnSessionDB[0] = sqlite.VpnSession{CrasIndex: int64(2000 + rand.Intn(100)),
			CrasUsername:        "test111111",
			CrasGroup:           "vpn",
			CrasSessionDuration: 1231,
			CrasLocalAddress:    "202.120.95.1",
			CrasISPAddress:      "2.123.123.123"}
		err := sqlite.Insert_DB(VpnSessionDB)
		sqlite.UpdateTime_DB(1101)
		sqlite.UpdateTime_DB(1102)
		sqlite.UpdateTime_DB(1103)
		sqlite.DeleteOld_DB(20)
		sessions, err := sqlite.SearchAll_DB()
		//sessions, err := sqlite.Search_DB("1101", 3)
		log.Println("第二个函数", sessions, err)
	})
	c.Start()
}
