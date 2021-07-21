package test

import (
	"os"
	"testing"
	"time"
)

func TestExit(t *testing.T) {

	go func() {
		time.Sleep(3*time.Second)
		println("exit")
		os.Exit(0)
	}()

	for  {
		time.Sleep(1*time.Second)
		println("main...")
	}
}
