package zset

import (
    "GoLearn/GoZset/list"
    "GoLearn/GoZset/node"
)

type ZSet struct {
    dict map[int64]int64
    Zsl  *list.ZSkipList
}

func (this *ZSet) Add(key,score int64)node.Node  {
    return this.Zsl.Add(key,score)
}
