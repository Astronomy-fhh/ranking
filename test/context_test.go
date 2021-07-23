package test

import (
	"testing"
	"time"
)

func TestContext(t *testing.T) {

	//ctx := context.Background()
	//
	//ctxc, _ := context.WithCancel(ctx)
	//
	//time.Sleep(10*time.Second)
	//ctx.Done()


	go func() {
		for  {
			time.Sleep(5*time.Second)
			println("go")
		}
	}()

	println("exit")
	//os.Exit(0)
}