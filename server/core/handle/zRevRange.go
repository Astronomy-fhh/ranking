package handle

import (
	"context"
	"ranking/db"
	"ranking/log"
	pb "ranking/proto"
)

func (ServerHandle) ZRevRange(ctx context.Context, req *pb.ZRevRangeReq) (*pb.ZRevRangeResp, error) {
	log.Log.Debugf("serverHandle:ZRevRange:req:%v", req)
	handle := &ZRevRangeHandle{req: req}
	err := serverHandle.Execute(ctx, handle)
	if err != nil {
		return nil, err
	}
	return handle.resp, nil
}

type ZRevRangeHandle struct {
	req  *pb.ZRevRangeReq
	resp *pb.ZRevRangeResp
}

func (h *ZRevRangeHandle) Validate() error { return nil }
func (h *ZRevRangeHandle) Execute() error {

	objs := db.Db.ZRevRange(h.req.Key, h.req.Start, h.req.End)
	h.resp = &pb.ZRevRangeResp{
		Objs: objs,
	}
	log.Log.Debugf("serverHandle:ZRevRange:resp:%v", h.resp)
	return nil
}
