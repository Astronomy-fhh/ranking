package main

import (
	"bufio"
	"fmt"
	"os"
	"ranking/list"
	"ranking/zset"
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
	for i := 0; i < 10; i++ {
		//n := rand.Int63n(10)
		zset.Add(int64(i),int64(i))
	}

	for i := 0; i < 5; i++ {
		//n := rand.Int63n(10)
		del := zset.Del(int64(i),int64(i))
		println(del)
	}

	op := zset.Zsl.Header
	for op.Layer[0].ForwardNode != nil  {
		score := op.Layer[0].ForwardNode.Score
		println(score)
		op = op.Layer[0].ForwardNode
	}
}

