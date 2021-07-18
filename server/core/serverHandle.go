package core

import (
	"context"
	"ranking/db"
	"ranking/log"
	pb "ranking/proto"
)

type ServerHandle struct {
	pb.UnimplementedRankServer
}

func NewServerHandle()*ServerHandle {
	return &ServerHandle{}
}


func (h *ServerHandle)ZAdd(ctx context.Context,req *pb.ZAddReq)(*pb.ZAddResp,error) {
	key := req.Key
	oneContainer := db.Db.GetOrInitContainer(key)
	c, u := oneContainer.Add(req.Val)
	res := &pb.ZAddResp{AddC: c,UpdateC: u}
	log.Log.Infof("serverHandle:ZAdd:ret:%v",res)
	return res,nil
}

func (h *ServerHandle)ZRem(ctx context.Context,req *pb.ZRemReq)(*pb.ZRemResp,error)  {
	return nil,nil
}

//func (h *ServerHandle)mustEmbedUnimplementedRankServer()  {
//}



