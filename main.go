package main

import (
	"GoLearn/GoZset/list"
	"GoLearn/GoZset/zset"
	"bufio"
	"fmt"
	"math/rand"
	"os"
)


func main() {
	container := initContainer()
	//initClient(container)
	test(container)
}

func initContainer()*zset.ZSet  {
	fmt.Println("initContainer...")

	zset := &zset.ZSet{
		Zsl: list.NewZSkipList(),
	}
	return zset
}

func initClient(zset *zset.ZSet)  {

	var key int64
	var score int64
	for  {
		fmt.Println("tab key：")
		stdin := bufio.NewReader(os.Stdin)
		val, err := fmt.Fscan(stdin, &key)
		if err != nil  {
			fmt.Println("tab err:",val)
			continue
		}

		fmt.Println("tab score：")
		val, err = fmt.Fscan(stdin, &score)
		if err != nil  {
			fmt.Println("tab err:",val)
			continue
		}

		addNode := zset.Add(key, score)
		fmt.Printf("%v", addNode)
		fmt.Println("ok")
	}
}

func test(zset *zset.ZSet)  {
	for i := 0; i < 10000; i++ {
		n := rand.Int63n(1000000)
		zset.Add(n,n)
	}

	op := zset.Zsl.Header
	for op.Layer[0].ForwardNode != nil  {
		score := op.Layer[0].ForwardNode.Score
		println(score)
		op = op.Layer[0].ForwardNode
	}
}

