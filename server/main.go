package main

import (
	"flag"
	"fmt"
	"os"
	"ranking/log"
	"ranking/server/config"
	"ranking/server/servermain"
)

var (
	conf  = flag.String("conf", "/Users/fanhuhu/PhpstormProjects/GOPATH/src/ranking/conf/server.json", "config file path")
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

	server := servermain.NewServer()
	server.Run()
}
