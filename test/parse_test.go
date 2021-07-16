package test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	a := " ZADD  2 a  3 b"
	b := strings.TrimSpace(a)
	println(b)
	compile := regexp.MustCompile("\\s+")
	split := compile.Split(b, -1)
	fmt.Printf("s%v",split)

	split = split[1:]
	fmt.Printf("s%v",split)

}









