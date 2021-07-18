package main

import (
	"flag"
	"fmt"
	"os"
	"ranking/config"
	"ranking/log"
	"ranking/server/core"
)

var (
	conf  = flag.String("conf", "/Users/huhu.fan/workspace/go/src/ranking/conf/server.json", "config file path")
	logProd  = flag.Bool("logProd", false, "production log")
)

func main()  {

	flag.Parse()

	if *conf == "" {
		flag.Usage()
		os.Exit(1)
	}


	err := log.Init(*logProd, "")
	if err != nil {
		fmt.Printf("init logger err:%s",err.Error())
		os.Exit(1)
	}

	err = config.InitServerConfig(*conf)
	if err != nil {
		fmt.Printf("get config err:%s",err.Error())
		os.Exit(1)
	}

	server := core.NewServer()
	server.Run()
}
