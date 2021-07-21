package handle

import (
	"context"
	"ranking/db"
	"ranking/log"
	pb "ranking/proto"
)

func (ServerHandle) ZIncrBy(ctx context.Context, req *pb.ZIncrByReq) (*pb.ZIncrByResp, error) {
	log.Log.Debugf("serverHandle:ZIncrBy:req:%v", req)
	handle := &ZIncrByHandle{req: req}
	err := serverHandle.Execute(ctx, handle)
	if err != nil {
		return nil, err
	}
	return handle.resp, nil
}

type ZIncrByHandle struct {
	req  *pb.ZIncrByReq
	resp *pb.ZIncrByResp
}

func (h *ZIncrByHandle) Validate() error { return nil }
func (h *ZIncrByHandle) Execute() error {

	ret := db.Db.ZIncrBy(h.req.Key,h.req.Incr,h.req.Member)
	h.resp = &pb.ZIncrByResp{
		Ret: &ret,
	}
	log.Log.Debugf("serverHandle:ZIncrBy:resp:%v", h.resp)
	return nil
}
