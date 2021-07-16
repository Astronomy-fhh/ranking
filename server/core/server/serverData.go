package server

import (
	"ranking/container"
	"ranking/log"
	"sync"
)

var serverData *ServerData

type ServerData struct {
	container map[string]*container.Container
	sync.Mutex
}

func ServerDataInit() {
	serverData = &ServerData{
		container: make(map[string]*container.Container),
	}
}

func (sd *ServerData) GetOrInitContainer(key string) *container.Container {
	_, ok := sd.container[key]
	if !ok {
		sd.Lock()
		defer sd.Unlock()
		_, ok := sd.container[key]
		if !ok {
			sd.container[key] = container.NewContainer()
			log.Log.Infof("init data container for key:%s", key)
		}
	}
	return sd.container[key]
}
