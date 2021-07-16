package handle


import (
	"google.golang.org/grpc"
	pb "ranking/proto"
)

var Handle *ClientHandle

type ClientHandle struct {
	ClientConn *grpc.ClientConn
	RankClient pb.RankClient
}

// distributed to different handles
func HandOut(method pb.Method,args []string)error{
	switch method {
	case pb.Method_ZADD:
		return Cmd(ZaddHandle(),args)
	case pb.Method_ZREM:
		return nil
	}
	return nil
}



type BaseHandleInterface interface {
	// parsing command  params
	Parse([]string)error
	// execute main logic
	Execute()error
	// command output
	Print()error
}


func Cmd(h BaseHandleInterface,args []string)error  {
	err := h.Parse(args)
	if err != nil {
		return err
	}
	err = h.Execute()
	if err != nil {
		return err
	}
	err = h.Print()
	return err
}