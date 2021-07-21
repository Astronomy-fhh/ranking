package db

import (
	pb "ranking/proto"
	"sync"
)

type Container struct {
	Dict map[string]int64
	Zsl  *ZSkipList
	sync.RWMutex
}

func NewContainer() *Container {
	zsl := NewZSkipList()
	dict := make(map[string]int64)
	return &Container{
		Dict: dict,
		Zsl:  zsl,
	}
}

func (c *Container) Size() int64 {
	return int64(len(c.Dict))
}

func (c *Container) Add(val map[string]int64) (int64, int64) {
	c.Lock()
	defer c.Unlock()
	var addC int64
	var updateC int64

	for member, score := range val {
		oldScore, ok := c.Dict[member]
		if ok {
			c.Zsl.Update(member, oldScore, score)
			updateC += 1
		} else {
			c.Zsl.Add(member, score)
			addC += 1
		}
		c.Dict[member] = score
	}
	return addC, updateC
}

func (c *Container) Inceby(incr int64, member string) int64 {
	c.Lock()
	defer c.Unlock()

	newScore := incr
	oldScore, ok := c.Dict[member]
	if ok {
		newScore = oldScore + newScore
		c.Zsl.Update(member, oldScore, newScore)
	} else {
		c.Zsl.Add(member, newScore)
	}

	c.Dict[member] = newScore
	return newScore
}

func (c *Container) DelMembers(members []string) int64 {
	c.Lock()
	defer c.Unlock()
	var ret int64
	for _, member := range members {
		score, ok := c.Dict[member]
		if !ok {
			continue
		}
		c.Zsl.Del(member, score)
		delete(c.Dict, member)
		ret++
	}
	return ret
}

func (c *Container) GetRangeByRank(start, end int64) []*pb.Obj {
	c.RLock()
	defer c.RUnlock()
	return c.Zsl.GetRangeByRank(start, end)
}

func (c *Container) GetRevRangeByRank(start, end int64) []*pb.Obj {
	c.RLock()
	defer c.RUnlock()
	return c.Zsl.GetRevRangeByRank(start, end)
}

func (c *Container) GetCountByRangeScore(start, end int64) int64 {
	c.RLock()
	defer c.RUnlock()
	return c.Zsl.GetCountByRangeScore(start, end)
}

func (c *Container) GetRevRank(member string) (int64, bool) {
	c.RLock()
	defer c.RUnlock()
	score, ok := c.Dict[member]
	if !ok {
		return 0, false
	}
	rank := c.Size() - c.Zsl.GetRank(member, score) - 1
	return rank, true
}

func (c *Container) GetScore(member string) (int64, bool) {
	c.RLock()
	defer c.RUnlock()
	score, ok := c.Dict[member]
	if !ok {
		return 0, false
	}
	return score, true
}

func (c *Container) GetRangeByScore(start, end int64) []*pb.Obj {
	c.RLock()
	defer c.RUnlock()
	return c.Zsl.GetRangeByScore(start, end)
}

func (c *Container) GetRevRangeByScore(start, end int64) []*pb.Obj {
	c.RLock()
	defer c.RUnlock()
	return c.Zsl.GetRevRangeByScore(start, end)
}

func (c *Container) GetRank(member string) (int64, bool) {
	c.RLock()
	defer c.RUnlock()
	score, ok := c.Dict[member]
	if !ok {
		return 0, false
	}
	rank := c.Zsl.GetRank(member, score)
	return rank, true
}

func (c *Container) GetLen() int64 {
	c.Lock()
	defer c.Unlock()
	return c.Zsl.Len
}

func (c *Container) DelRangeByRank(start, end int64) int64 {
	c.Lock()
	defer c.Unlock()
	var ret int64
	objs := c.Zsl.GetRangeByRank(start, end)
	for _, obj := range objs {
		score, ok := c.Dict[obj.Member]
		if !ok {
			continue
		}
		c.Zsl.Del(obj.Member, score)
		delete(c.Dict, obj.Member)
		ret++
	}
	return ret
}

func (c *Container) DelRangeByScore(start, end int64) int64 {
	c.Lock()
	defer c.Unlock()
	var ret int64
	objs := c.Zsl.GetRangeByScore(start, end)
	for _, obj := range objs {
		score, ok := c.Dict[obj.Member]
		if !ok {
			continue
		}
		c.Zsl.Del(obj.Member, score)
		delete(c.Dict, obj.Member)
		ret++
	}
	return ret
}
