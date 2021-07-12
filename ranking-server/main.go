package main

import (
	"fmt"
	"ranking/container"
	"strconv"
	"time"
)


func main() {
	container := initContainer()
	test(container)
}

func initContainer()*container.ZSet {
	fmt.Println("initContainer...")

	zset := &container.ZSet{
		Zsl: container.NewZSkipList(),
	}
	return zset
}

func test(zset *container.ZSet)  {
	for i := 1; i <= 100000; i++ {
		//n := rand.Int63n(10)
		go zset.Add(strconv.Itoa(i),int64(i))
	}
	for  {
		time.Sleep(100000000)
		fmt.Printf("%d\n",zset.GetLen())
	}
}

