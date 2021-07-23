package db

import (
	"bufio"
	"encoding/binary"
	"go.uber.org/atomic"
	"google.golang.org/protobuf/proto"
	"io"
	"math"
	"math/rand"
	"os"
	"ranking/log"
	pb "ranking/proto"
	"ranking/server/config"
	"ranking/util"
	"sync"
	"syscall"
	"time"
)

var Db *DB

type DB struct {
	version    uint64
	initTime   uint64
	stop  atomic.Bool
	containers map[string]*Container
	sync.RWMutex
}

func InitDB() {
	Db = &DB{
		containers: make(map[string]*Container),
	}

	err := Db.RDBLoad()
	if err != nil {
		log.Log.Fatalf("RDBLoad:fail:%v", err.Error())
	}
	Db.GenTestData()
	go Db.RDBSaveLoop()
}

func (db *DB) RDBSaveLoop() {
	timeInterval := config.SConfig.RDBTimeIntervals
	log.Log.Infof("RDBSaveLoop:start:timeInterval:%v", timeInterval)

	ticker := time.NewTicker(time.Duration(timeInterval) * time.Second)

	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			db.RDBSave()
		}
	}
}

func (db *DB) RDBSave() {
	db.Lock()
	defer db.Unlock()
	log.Log.Info("RDBSave:start...")
	rdb := &pb.RDB{}
	data := make(map[string]*pb.Container)
	rdb.Version = db.version + 1
	for key, container := range db.containers {
		container.RLock()
		dict := make(map[string]int64,len(container.Dict))
		for member, score := range container.Dict {
			dict[member] = score
		}
		container.RUnlock()
		c := pb.Container{
			Data: dict,
		}
		data[key] = &c
	}
	rdb.Containers = data
	bytes, err := proto.Marshal(rdb)
	if err != nil {
		log.Log.Warnf("RDBSave:fail:%v", err.Error())
		return
	}
	head := make([]byte, 8)
	binary.BigEndian.PutUint64(head, uint64(len(bytes)))
	log.Log.Debugf("RDBSave:size:%v", uint64(len(bytes)))

	bytes = append(head, bytes...)
	err = saveBytes(bytes, config.SConfig.RDBFileName)
	if err != nil {
		log.Log.Warnf("RDBSave:fail:%v", err.Error())
		return
	}
	log.Log.Info("RDBSave:ok")
}

func (db *DB) RDBLoad() error {
	log.Log.Info("RDBLoad:start...")

	err := syscall.Access(config.SConfig.RDBFileName, syscall.F_OK)
	if err != nil {
		log.Log.Warnf("RDBLoad:not found file:%v", config.SConfig.RDBFileName)
		return nil
	}

	f, err := os.OpenFile(config.SConfig.RDBFileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(f)

	var payloadSize uint64
	err = binary.Read(reader, binary.BigEndian, &payloadSize)
	if err != nil {
		return err
	}
	log.Log.Debugf("RDBLoad:payloadSize:%d", payloadSize)

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

func saveBytes(bytes []byte, filename string) error {
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
	db.RLock()
	_, ok := db.containers[key]
	db.RUnlock()
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

func (db *DB) GetContainer(key string) *Container {
	db.RLock()
	defer db.RUnlock()
	_, ok := db.containers[key]
	if !ok {
		return nil
	}
	return db.containers[key]
}

func (db *DB) GenTestData() {
	log.Log.Info("GenTestData...start")

	keyC := 1000
	keyL := 10
	mC := 1000
	mL := 15
	sL := 2

	runLimit := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		runLimit <- true
	}
	for i := 0; i < keyC; i++ {
		select {
		case <-runLimit:
			go func(i int) {
				key := util.GetRandomString(keyL)
				c := db.GetOrInitContainer(key)
				val := make(map[string]int64, mC)
				for j := 0; j < mC; j++ {
					member := util.GetRandomString(mL)
					score := rand.Int63n(int64(math.Pow(10, float64(sL))))
					val[member] = score
				}
				a, u := c.Add(val)
				log.Log.Infof("GenTestData:key:%d:%s,addC:%d,updateC:%d", i, key, a, u)
				runLimit <- true
			}(i)
		}
	}
}

func (db *DB) AllObjs() map[string][]*pb.Obj {
	payload := make(map[string][]*pb.Obj)
	for key, container := range db.containers {
		ranks := container.GetRangeByRank(0, -1)
		payload[key] = ranks
	}
	return payload
}

func (db *DB) ZRem(key string, members []string) int64 {
	var ret int64
	c := db.GetContainer(key)
	if c == nil {
		return ret
	}
	return c.DelMembers(members)
}

func (db *DB) ZRemRangeByRank(key string, start, end int64) int64 {
	var ret int64
	c := db.GetContainer(key)
	if c == nil {
		return ret
	}
	return c.DelRangeByRank(start, end)
}

func (db *DB) ZRemRangeByScore(key string, start int64, end int64) int64 {
	var ret int64
	c := db.GetContainer(key)
	if c == nil {
		return ret
	}
	return c.DelRangeByScore(start, end)
}

func (db *DB) ZRevRange(key string, start int64, end int64) []*pb.Obj {
	c := db.GetContainer(key)
	if c == nil {
		ret := make([]*pb.Obj, 0)
		return ret
	}
	objs := c.GetRevRangeByRank(start, end)
	return objs
}

func (db *DB) ZRevRangeByScore(key string, start int64, end int64) []*pb.Obj {
	c := db.GetContainer(key)
	if c == nil {
		res := make([]*pb.Obj, 0)
		return res
	}
	objs := c.GetRevRangeByScore(start, end)
	return objs
}

func (db *DB) ZRevRank(key string, member string) (int64, bool) {
	var ret int64
	c := db.GetContainer(key)
	if c == nil {
		return ret, false
	}
	return c.GetRevRank(member)
}

func (db *DB) ZScore(key string, member string) (int64, bool) {
	var ret int64
	c := db.GetContainer(key)
	if c == nil {
		return ret, false
	}
	return c.GetScore(member)
}

func (db *DB) ZCard(key string) (int64, bool) {
	var ret int64
	c := db.GetContainer(key)
	if c == nil {
		return ret, false
	}
	c.RLock()
	defer c.RUnlock()
	return c.Size(), true
}

func (db *DB) ZAdd(key string, vars map[string]int64) (int64, int64) {
	c := db.GetOrInitContainer(key)
	return c.Add(vars)
}

func (db *DB) ZCount(key string, start, end int64) (int64, bool) {
	var ret int64
	c := db.GetContainer(key)
	if c == nil {
		return ret, false
	} else {
		return c.GetCountByRangeScore(start, end), true
	}
}

func (db *DB) ZIncrBy(key string, incr int64, member string) int64 {
	container := db.GetOrInitContainer(key)
	return container.Inceby(incr, member)
}

func (db *DB) ZRange(key string, start int64, end int64) []*pb.Obj {
	ret := make([]*pb.Obj, 0)
	c := db.GetContainer(key)
	if c != nil {
		ret = c.GetRangeByRank(start, end)
	}
	return ret
}

func (db *DB) ZRangeByScore(key string, start int64, end int64) []*pb.Obj {
	ret := make([]*pb.Obj, 0)
	c := db.GetContainer(key)
	if c != nil {
		ret = c.GetRangeByScore(start, end)
	}
	return ret
}

func (db *DB) ZRank(key string, member string) (int64, bool) {
	c := db.GetContainer(key)
	if c != nil {
		return c.GetRank(member)
	}
	return 0, false
}
