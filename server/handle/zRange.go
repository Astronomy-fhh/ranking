package handle

import (
	"context"
	"ranking/db"
	"ranking/log"
	pb "ranking/proto"
)

func (ServerHandle) ZRange(ctx context.Context, req *pb.ZRangeReq) (*pb.ZRangeResp, error) {
	log.Log.Debugf("serverHandle:ZRange:req:%v", req)
	handle := &ZRangeHandle{req: req}
	err := serverHandle.Execute(ctx, handle)
	if err != nil {
		return nil, err
	}
	return handle.resp, nil
}

type ZRangeHandle struct {
	req  *pb.ZRangeReq
	resp *pb.ZRangeResp
}

func (h *ZRangeHandle) Validate() error { return nil }
func (h *ZRangeHandle) Execute() error {
	ret := db.Db.ZRange(h.req.Key, h.req.Start, h.req.End)
	h.resp = &pb.ZRangeResp{
		Objs: ret,
	}
	log.Log.Debugf("serverHandle:ZRange:resp:%v", h.resp)
	return nil
}
