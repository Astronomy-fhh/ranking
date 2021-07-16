package container

import (
    "ranking/log"
    "sync"
)

type Container struct {
    Dict map[string]uint64
    Zsl  *ZSkipList
    rwMutex sync.RWMutex
}


func NewContainer()*Container  {
    zsl := NewZSkipList()
    dict := make(map[string]uint64)
    return &Container{
        Dict: dict,
        Zsl: zsl,
    }
}


func (zst *Container) Add(val map[string]uint64)(uint64,uint64) {
    log.Log.Infof("container:Add:%v",val)
    zst.rwMutex.Lock()
    defer zst.rwMutex.Unlock()
    var addC uint64 = 0
    var updateC uint64 = 0

    for member, score := range val {
        oldScore,ok := zst.Dict[member]
        if ok {
            zst.Zsl.Update(member,oldScore,score)
            updateC += 1
        }else{
            zst.Zsl.Add(member,score)
            addC += 1
        }
        zst.Dict[member] = score
    }
    return addC,updateC
}

func (zst *Container) Del(key string, score uint64)bool  {
    zst.rwMutex.Lock()
    defer zst.rwMutex.Unlock()
    return zst.Zsl.Del(key,score)
}

func (zst *Container) GetRangeByRank(min uint64, max int64)[]Obj {
    zst.rwMutex.RLock()
    defer zst.rwMutex.RUnlock()
    return zst.Zsl.GetRangeByRank(min,max)
}

func (zst *Container) GetRevRangeByRank(min uint64, max int64)[]Obj {
    zst.rwMutex.RLock()
    defer zst.rwMutex.RUnlock()
    return zst.Zsl.GetRevRangeByRank(min,max)
}

func (zst *Container) GetRangeByScore(min uint64, max int64)[]Obj {
    zst.rwMutex.RLock()
    defer zst.rwMutex.RUnlock()
    return zst.Zsl.GetRangeByScore(min,max)
}

func (zst *Container) GetRevRangeByScore(min uint64, max int64)[]Obj {
    zst.rwMutex.RLock()
    defer zst.rwMutex.RUnlock()
    return zst.Zsl.GetRevRangeByScore(min,max)
}

func (zst *Container) GetLen()uint64{
    zst.rwMutex.Lock()
    defer zst.rwMutex.Unlock()
    return zst.Zsl.Len
}