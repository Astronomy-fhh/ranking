package handle

import (
	"context"
	"ranking/db"
	"ranking/log"
	pb "ranking/proto"
)

func (ServerHandle) ZScore(ctx context.Context, req *pb.ZScoreReq) (*pb.ZScoreResp, error) {
	log.Log.Debugf("serverHandle:ZScore:req:%v", req)
	handle := &ZScoreHandle{req: req}
	err := serverHandle.Execute(ctx, handle)
	if err != nil {
		return nil, err
	}
	return handle.resp, nil
}

type ZScoreHandle struct {
	req  *pb.ZScoreReq
	resp *pb.ZScoreResp
}

func (h *ZScoreHandle) Validate() error { return nil }

func (h *ZScoreHandle) Execute() error {
	resp := &pb.ZScoreResp{}
	score, has := db.Db.ZScore(h.req.Key, h.req.Member)
	if has {
		resp.Score = &score
	}
	h.resp = resp
	log.Log.Debugf("serverHandle:ZScore:resp:%v", h.resp)
	return nil
}
