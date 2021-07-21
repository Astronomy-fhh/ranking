package config

import (
	"encoding/json"
	"os"
	"ranking/log"
)


var CConfig *ClientConfig

type ClientConfig struct {
	HttpAddr        string
}


func InitClientConfig(fileName string)error  {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	decode := json.NewDecoder(file)
	config := &ClientConfig{}
	err = decode.Decode(config)
	if err != nil {
		return err
	}
	log.Log.Infof("configFromFile:%v", config)
	CConfig = config
	return nil
}

