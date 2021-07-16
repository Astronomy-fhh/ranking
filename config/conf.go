package config

import (
	"encoding/json"
	"os"
	"ranking/log"
)

type ServerConfig struct {
	HttpAddr        string
	ListMaxLayer    int32
	ListLayerFactor float32
}

type ClientConfig struct {
	HttpAddr        string
}


func NewServerConf(fileName string)(*ServerConfig,error)  {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decode := json.NewDecoder(file)
	config := &ServerConfig{}
	err = decode.Decode(config)
	if err != nil {
		return nil, err
	}
	log.Log.Infof("configFromFile:%v", config)
	return config, nil
}

func NewClientConf(fileName string)(*ClientConfig,error)  {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decode := json.NewDecoder(file)
	config := &ClientConfig{}
	err = decode.Decode(config)
	if err != nil {
		return nil, err
	}
	log.Log.Infof("configFromFile:%v", config)
	return config, nil
}

