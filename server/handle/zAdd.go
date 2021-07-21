package handle

import (
	"context"
	"ranking/db"
	"ranking/log"
	pb "ranking/proto"
	config2 "ranking/server/config"
	serverError "ranking/server/error"
	"strings"
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
	if strings.TrimSpace(h.req.Key) == "" {

	}
	if int64(len(vars)) > config2.SConfig.SingleZAddLimit {
		return serverError.NewInvalidArgumentError("many to add")
	}
	return nil
}

func (h *ZAddHandle) Execute()error  {
	add, update := db.Db.ZAdd(h.req.Key, h.req.Vars)
	h.resp = &pb.ZAddResp{AddC: &add,UpdateC: &update}
	log.Log.Debugf("serverHandle:ZAdd:resp:%v",h.resp)
	return nil
}

