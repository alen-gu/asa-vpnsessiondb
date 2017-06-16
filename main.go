package main

import (
	"flag"
	"fmt"
	"os"

	//	"github.com/alen-gu/asa-vpnsessiondb/cron"
	"github.com/alen-gu/asa-vpnsessiondb/g"

	"github.com/alen-gu/asa-vpnsessiondb/http"
)

func main() {

	cfg := flag.String("c", "cfg.example.json", "configuration file")
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)
	//sqlite.Init_DB()
	//go cron.UpdateDB()
	//aaa, err := sqlite.SearchAll_DB()
	//fmt.Println(err)
	//fmt.Println(aaa)
	go http.Start()
	select {}

}
