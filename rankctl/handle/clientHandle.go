package handle

import (
	"google.golang.org/grpc"
	"ranking/log"
	pb "ranking/proto"
)

var Handle *ClientHandle

type ClientHandle struct {
	ClientConn *grpc.ClientConn
	RankClient pb.RankClient
}

type BaseHandleInterface interface {
	// parsing command  params
	Parse([]string)error
	// execute main logic
	Execute()error
	// command output
	Print()error
}


func CmdHandle(h BaseHandleInterface,args []string)error  {
	err := h.Parse(args)
	if err != nil {
		return err
	}
	err = h.Execute()
	if err != nil {
		log.Log.Infof("cmd err:%v",err)
		return err
	}
	err = h.Print()
	return err
}
