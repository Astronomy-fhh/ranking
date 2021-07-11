package zset

import (
    "ranking/list"
    "ranking/node"
)

type ZSet struct {
    dict map[int64]int64
    Zsl  *list.ZSkipList
}

func (this *ZSet) Add(key,score int64)node.Node  {
    return this.Zsl.Add(key,score)
}


func (this *ZSet) Del(key, score int64)bool  {
    return this.Zsl.Del(key,score)
}

func (this *ZSet) GraphZsl()interface{}  {
    return this.Zsl.GraphPrint()
}