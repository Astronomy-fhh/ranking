package test

import (
	"fmt"
	"testing"
)

func TestPoint(t *testing.T) {

	a := make(map[string]int)
	a["ss"] = 1

	b := a

	fmt.Printf("a:%p",a)
	fmt.Println()
	fmt.Printf("b:%p",b)

}
