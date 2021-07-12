package main

import (
	"fmt"
	"ranking/list"
	"ranking/zset"
)


func main() {
	container := initContainer()
	test(container)
}

func initContainer()*zset.ZSet  {
	fmt.Println("initContainer...")

	zset := &zset.ZSet{
		Zsl: list.NewZSkipList(),
	}
	return zset
}

func test(zset *zset.ZSet)  {
	for i := 1; i <= 10000; i++ {
		//n := rand.Int63n(10)
		 zset.Add(int64(i),int64(i))
	}
}

