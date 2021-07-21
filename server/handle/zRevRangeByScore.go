package handle

import (
	"context"
	"ranking/db"
	"ranking/log"
	pb "ranking/proto"
)

func (ServerHandle) ZRevRangeByScore(ctx context.Context, req *pb.ZRevRangeByScoreReq) (*pb.ZRevRangeByScoreResp, error) {
	log.Log.Debugf("serverHandle:ZRevRangeByScore:req:%v", req)
	handle := &ZRevRangeByScoreHandle{req: req}
	err := serverHandle.Execute(ctx, handle)
	if err != nil {
		return nil, err
	}
	return handle.resp, nil
}

type ZRevRangeByScoreHandle struct {
	req  *pb.ZRevRangeByScoreReq
	resp *pb.ZRevRangeByScoreResp
}

func (h *ZRevRangeByScoreHandle) Validate() error { return nil }
func (h *ZRevRangeByScoreHandle) Execute() error {
	objs := db.Db.ZRevRangeByScore(h.req.Key, h.req.Start, h.req.End)
	h.resp = &pb.ZRevRangeByScoreResp{
		Objs: objs,
	}
	log.Log.Debugf("serverHandle:ZRevRangeByScore:resp:%v", h.resp)
	return nil
}
