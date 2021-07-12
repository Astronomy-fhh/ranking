package container

import (
    "sync"
)

type ZSet struct {
    dict *map[int64]int64
    Zsl  *ZSkipList
    rwMutex sync.RWMutex
}

func (zst *ZSet) Add(key string,score int64) Obj {
    zst.rwMutex.Lock()
    defer zst.rwMutex.Unlock()
    return zst.Zsl.Add(key,score)
}

func (zst *ZSet) Del(key string, score int64)bool  {
    zst.rwMutex.Lock()
    defer zst.rwMutex.Unlock()
    return zst.Zsl.Del(key,score)
}

func (zst *ZSet) GetRangeByRank(min, max int64)[]Obj {
    zst.rwMutex.RLock()
    defer zst.rwMutex.RUnlock()
    return zst.Zsl.GetRangeByRank(min,max)
}

func (zst *ZSet) GetRevRangeByRank(min, max int64)[]Obj {
    zst.rwMutex.RLock()
    defer zst.rwMutex.RUnlock()
    return zst.Zsl.GetRevRangeByRank(min,max)
}

func (zst *ZSet) GetRangeByScore(min, max int64)[]Obj {
    zst.rwMutex.RLock()
    defer zst.rwMutex.RUnlock()
    return zst.Zsl.GetRangeByScore(min,max)
}

func (zst *ZSet) GetRevRangeByScore(min, max int64)[]Obj {
    zst.rwMutex.RLock()
    defer zst.rwMutex.RUnlock()
    return zst.Zsl.GetRevRangeByScore(min,max)
}

func (zst *ZSet) GetLen()int64{
    zst.rwMutex.Lock()
    defer zst.rwMutex.Unlock()
    return zst.Zsl.Len
}