package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"ranking/list"
	"ranking/zset"
	"runtime"
)


func main() {
	container := initContainer()
	test(container)
	startGraphServer(container)
}

func startGraphServer(zset *zset.ZSet)  {
	http.HandleFunc("/zSkipListGraph", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		writer.Header().Set("content-type", "application/json")
		marshal, err := json.Marshal(zset.GraphZsl())
		if err != nil {
			_, _ = writer.Write([]byte(err.Error()))
		}
		_, _ = writer.Write(marshal)
	})
	_ = http.ListenAndServe(":8087", nil)
}


func initContainer()*zset.ZSet  {
	fmt.Println("initContainer...")

	zset := &zset.ZSet{
		Zsl: list.NewZSkipList(),
	}
	return zset
}

func test(zset *zset.ZSet)  {
	for i := 0; i < 100000; i++ {
		//n := rand.Int63n(10)
		zset.Add(int64(rand.Intn(100000000)),int64(rand.Intn(100000000)))
	}
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf("Alloc:%d(Mb) HeapIdle:%d(MB) HeapReleased:%d(MB)", ms.Alloc / 1024, ms.HeapIdle /1024, ms.HeapReleased/1024)


	//for i := 0; i < 5; i++ {
	//	//n := rand.Int63n(10)
	//	del := zset.Del(int64(i),int64(i))
	//	println(del)
	//}

	//zsl := zset.GraphZsl()
	//fmt.Printf("%v",zsl)
}

