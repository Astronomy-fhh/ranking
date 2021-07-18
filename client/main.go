package main

import (
	"fmt"
	"os"
	"ranking/client/core/client"
	"ranking/config"
	"ranking/log"
)

func main()  {

	//var (
	//	h  = flag.String("h", "", "host")
	//	p  = flag.String("p", "", "port")
	//	conf  = flag.String("conf", "", "config file path")
	//)

	logPath := "/Users/fanhuhu/PhpstormProjects/GOPATH/src/ranking/log/rank.log"
	confPath := "/Users/fanhuhu/PhpstormProjects/GOPATH/src/ranking/conf/client.json"
	err := log.Init(false,logPath)
	if err != nil {
		fmt.Printf("init logger err:%s",err.Error())
		os.Exit(1)
	}

	err = config.InitClientConfig(confPath)
	if err != nil {
		fmt.Printf("init config err:%s",err.Error())
		os.Exit(1)
	}

	client := client.NewClient()
	client.Run()

}

