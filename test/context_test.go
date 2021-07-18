package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {

	ctx := context.Background()

	ctxc, _ := context.WithCancel(ctx)

	time.Sleep(10*time.Second)
	ctx.Done()


	go func() {
		for  {
			select {
			case <-ctxc.Done():
				fmt.Printf("done:%d\n",1)
			default:
				fmt.Printf("doing...:%d\n",1)
			}
		}
	}()

}