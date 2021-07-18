package db

import (
	"bufio"
	"encoding/binary"
	"google.golang.org/protobuf/proto"
	"io"
	"math"
	"math/rand"
	"os"
	"ranking/config"
	"ranking/log"
	pb "ranking/proto"
	"ranking/util"
	"sync"
	"syscall"
	"time"
)

var Db *DB

type DB struct {
	version uint64
	initTime uint64
	containers map[string]*Container
	sync.Mutex
}

func InitDB() {
	Db = &DB{
		containers: make(map[string]*Container),
	}

	//err := Db.RDBLoad()
	//if err != nil {
	//	log.Log.Fatalf("RDBLoad:fail:%v",err.Error())
	//}

	Db.GenTestData()
	//go Db.RDBSaveLoop()
}

func (db *DB) RDBSaveLoop()  {
	timeInterval := config.SConfig.RDBTimeIntervals
	log.Log.Infof("RDBSaveLoop:start:timeInterval:%v",timeInterval)

	ticker := time.NewTicker(time.Duration(timeInterval) * time.Second)
	canSaveChan := make(chan bool ,1)
	canSaveChan <- true
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			select {
			case <-canSaveChan :
				db.RDBSave()
				canSaveChan <- true
			}
		}
	}
}

func (db *DB) RDBSave()  {
	log.Log.Info("RDBSave:start...")
	db.Lock()
	defer db.Unlock()
	rdb := &pb.RDB{}
	data := make(map[string]*pb.Container)
	rdb.Version = db.version + 1
	for key, container := range db.containers {
		container.Lock()
		c := pb.Container{
			Data: container.Dict,
		}
		data[key] = &c
		container.Unlock()
	}
	rdb.Containers = data
	bytes, err := proto.Marshal(rdb)
	if err != nil {
		log.Log.Warnf("RDBSave:fail:%v",err.Error())
		return
	}
	head := make([]byte,8)
	binary.BigEndian.PutUint64(head, uint64(len(bytes)))
	log.Log.Debugf("RDBSave:size:%v",uint64(len(bytes)))

	bytes = append(head,bytes...)
	err = saveBytes(bytes, config.SConfig.RDBFileName)
	if err != nil {
		log.Log.Warnf("RDBSave:fail:%v",err.Error())
		return
	}
	log.Log.Info("RDBSave:ok")
}

func (db *DB) RDBLoad()error  {
	log.Log.Info("RDBLoad:start...")

	err := syscall.Access(config.SConfig.RDBFileName, syscall.F_OK)
	if err != nil {
		log.Log.Warnf("RDBLoad:not found file:%v",config.SConfig.RDBFileName)
		return nil
	}

	f, err := os.OpenFile(config.SConfig.RDBFileName, os.O_RDONLY , os.ModePerm)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(f)

	var payloadSize uint64
	err = binary.Read(reader, binary.BigEndian, &payloadSize)
	if err != nil {
		return err
	}
	log.Log.Debugf("RDBLoad:payloadSize:%d",payloadSize)

	payload := make([]byte, payloadSize)
	err = binary.Read(reader, binary.BigEndian, &payload)
	if err != nil {
		return err
	}
	rdb := &pb.RDB{}
	err = proto.Unmarshal(payload, rdb)
	if err != nil {
		return err
	}
    Db.version = rdb.Version
    Db.initTime = rdb.Timestamp
	for key, containerData := range rdb.Containers {
		container := db.GetOrInitContainer(key)
		dict := containerData.Data
		container.Add(dict)
	}
	log.Log.Info("RDBLoad:ok")
	return nil
}

func saveBytes(bytes []byte,filename string)error  {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	n, err := f.Write(bytes)
	if err == nil && n < len(bytes) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

func (db *DB) GetOrInitContainer(key string) *Container {
	_, ok := db.containers[key]
	if !ok {
		db.Lock()
		defer db.Unlock()
		_, ok := db.containers[key]
		if !ok {
			db.containers[key] = NewContainer()
			log.Log.Infof("init container for key:%s", key)
		}
	}
	return db.containers[key]
}

func (db *DB) GenTestData()  {
	keyC := 1000
	keyL := 15
	mC := 10000
	mL := 15
	sL := 10

	for i := 0; i < keyC; i++ {
			key := util.GetRandomString(keyL)
			container := db.GetOrInitContainer(key)
			val := make(map[string]uint64,mC)
			for j := 0; j < mC; j++ {
				member := util.GetRandomString(mL)
				score := rand.Int63n(int64(math.Pow(10, float64(sL))))
				val[member] = uint64(score)
			}
			a, u := container.Add(val)
			log.Log.Infof("GenTestData:key:%d:%s,addC:%d,updateC:%d", i,key,a,u)
	}
}

func (db *DB) AllObjs()map[string][]*Obj  {
	payload := make(map[string][]*Obj)
	for key, container := range db.containers {
		ranks := container.GetRangeByRank(0, -1)
		payload[key] = ranks
	}
	return payload
}
