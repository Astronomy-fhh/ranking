package container

import (
    "sync"
)

type Container struct {
    dict *map[int64]int64
    Zsl  *ZSkipList
    rwMutex sync.RWMutex
}

func NewContainer()*Container  {
    zsl := NewZSkipList()
    return &Container{
        Zsl: zsl,
    }
}

func (zst *Container) Add(key string,score int64) Obj {
    zst.rwMutex.Lock()
    defer zst.rwMutex.Unlock()
    return zst.Zsl.Add(key,score)
}

func (zst *Container) Del(key string, score int64)bool  {
    zst.rwMutex.Lock()
    defer zst.rwMutex.Unlock()
    return zst.Zsl.Del(key,score)
}

func (zst *Container) GetRangeByRank(min, max int64)[]Obj {
    zst.rwMutex.RLock()
    defer zst.rwMutex.RUnlock()
    return zst.Zsl.GetRangeByRank(min,max)
}

func (zst *Container) GetRevRangeByRank(min, max int64)[]Obj {
    zst.rwMutex.RLock()
    defer zst.rwMutex.RUnlock()
    return zst.Zsl.GetRevRangeByRank(min,max)
}

func (zst *Container) GetRangeByScore(min, max int64)[]Obj {
    zst.rwMutex.RLock()
    defer zst.rwMutex.RUnlock()
    return zst.Zsl.GetRangeByScore(min,max)
}

func (zst *Container) GetRevRangeByScore(min, max int64)[]Obj {
    zst.rwMutex.RLock()
    defer zst.rwMutex.RUnlock()
    return zst.Zsl.GetRevRangeByScore(min,max)
}

func (zst *Container) GetLen()int64{
    zst.rwMutex.Lock()
    defer zst.rwMutex.Unlock()
    return zst.Zsl.Len
}