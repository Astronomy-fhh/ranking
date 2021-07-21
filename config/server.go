package config

import (
	"encoding/json"
	"os"
	"ranking/log"
)


var SConfig *ServerConfig

type ServerConfig struct {
	HttpAddr        string
	StatusHttpAddr string
	ListMaxLayer    int64
	ListLayerFactor float32
	RDBTimeIntervals int64
	RDBModifyKeys int64
	RDBFileName string
	SingleZAddLimit int64
}


func InitServerConfig(fileName string)error  {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	decode := json.NewDecoder(file)
	config := &ServerConfig{}
	err = decode.Decode(config)
	if err != nil {
		return err
	}
	log.Log.Infof("configFromFile:%v", config)
	SConfig = config
	return nil
}

