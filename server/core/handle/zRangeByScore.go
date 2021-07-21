package handle

import (
	"context"
	"ranking/db"
	"ranking/log"
	pb "ranking/proto"
)

func (ServerHandle) ZRangeByScore(ctx context.Context, req *pb.ZRangeByScoreReq) (*pb.ZRangeByScoreResp, error) {
	log.Log.Debugf("serverHandle:ZRangeByScore:req:%v", req)
	handle := &ZRangeByScoreHandle{req: req}
	err := serverHandle.Execute(ctx, handle)
	if err != nil {
		return nil, err
	}
	return handle.resp, nil
}

type ZRangeByScoreHandle struct {
	req  *pb.ZRangeByScoreReq
	resp *pb.ZRangeByScoreResp
}

func (h *ZRangeByScoreHandle) Validate() error { return nil }
func (h *ZRangeByScoreHandle) Execute() error {

	objs := db.Db.ZRangeByScore(h.req.Key, h.req.Start, h.req.End)

	h.resp = &pb.ZRangeByScoreResp{
		Objs: objs,
	}
	log.Log.Debugf("serverHandle:ZRangeByScore:resp:%v", h.resp)
	return nil
}
