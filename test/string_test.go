package test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestTrim(t *testing.T) {
	a := " as dd ss   "
	b := strings.TrimSpace(a)
	println(b)
	compile := regexp.MustCompile("\\s+")
	split := compile.Split(b, -1)
	fmt.Printf("s%v",split)

}









