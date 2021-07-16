package test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	var i uint32 = 1234
	fmt.Printf("%02X ", i)
	fmt.Printf("\n")

	// 小端
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, i)
	fmt.Printf("LittleEndian(%d) :", i)
	for _, bin := range b {
		fmt.Printf("%02X ", bin)
	}
	fmt.Printf("\n")

	//大端
	fmt.Printf("BigEndian(%d) :", i)
	binary.BigEndian.PutUint32(b, i)
	for _, bin := range b {
		fmt.Printf("%02X ", bin)
	}
	fmt.Printf("\n")

	//[]byte 2 uint32
	bytesBuffer := bytes.NewBuffer(b)
	var j uint32
	binary.Read(bytesBuffer, binary.BigEndian, &j)
	fmt.Println("j = ", j)
}
