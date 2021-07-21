package handle

import (
	"context"
	pb "ranking/proto"
)

var serverHandle *ServerHandle

type ServerHandle struct {
	pb.UnimplementedRankServer
}

func NewServerHandle()*ServerHandle {
	serverHandle = &ServerHandle{}
	return serverHandle
}

type BaseHandleInterface interface {
	Validate()error
	Execute()error
}


func (ServerHandle) Execute(ctx context.Context, h BaseHandleInterface)error{
	err := h.Validate()
	if err != nil {
		return err
	}
	err = h.Execute()
	if err != nil {
		return err
	}
	return nil
}