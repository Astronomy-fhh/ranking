package util

import (
	"math/rand"
	"time"
)


func GetRandomString(max int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()+rand.Int63n(1000)))
	for i := 0; i < max; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
