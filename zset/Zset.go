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

func (this *ZSet) GetRangeByRank(min, max int64)[]node.Node{
    return this.Zsl.GetRangeByRank(min,max)
}

func (this *ZSet) GetRevRangeByRank(min, max int64)[]node.Node{
    return this.Zsl.GetRevRangeByRank(min,max)
}

func (this *ZSet) GetRangeByScore(min, max int64)[]node.Node{
    return this.Zsl.GetRangeByScore(min,max)
}

func (this *ZSet) GetRevRangeByScore(min, max int64)[]node.Node{
    return this.Zsl.GetRevRangeByScore(min,max)
}

