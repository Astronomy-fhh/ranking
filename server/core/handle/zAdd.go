package handle

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ranking/config"
	"ranking/db"
	"ranking/log"
	pb "ranking/proto"
)

func (ServerHandle)ZAdd(ctx context.Context,req *pb.ZAddReq)(*pb.ZAddResp,error) {
	log.Log.Debugf("serverHandle:ZAdd:req:%v",req)

	handle := &ZAddHandle{req: req}
	err := serverHandle.Execute(ctx, handle)
	if err != nil {
		return nil, err
    }
	return handle.resp,nil
}

type ZAddHandle struct {
	req *pb.ZAddReq
	resp *pb.ZAddResp
}

func (h *ZAddHandle) Validate()error  {
	vars := h.req.Vars
	if int64(len(vars)) > config.SConfig.SingleZAddLimit {
		s := status.New(codes.OutOfRange, "many to add")
		return s.Err()
	}
	return nil
}

func (h *ZAddHandle) Execute()error  {
	add, update := db.Db.ZAdd(h.req.Key, h.req.Vars)
	h.resp = &pb.ZAddResp{AddC: &add,UpdateC: &update}
	log.Log.Debugf("serverHandle:ZAdd:resp:%v",h.resp)
	return nil
}

