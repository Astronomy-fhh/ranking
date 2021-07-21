package handle

import (
	"context"
	"ranking/db"
	"ranking/log"
	pb "ranking/proto"
)

func (ServerHandle) ZRevRank(ctx context.Context, req *pb.ZRevRankReq) (*pb.ZRevRankResp, error) {
	log.Log.Debugf("serverHandle:ZRevRank:req:%v", req)
	handle := &ZRevRankHandle{req: req}
	err := serverHandle.Execute(ctx, handle)
	if err != nil {
		return nil, err
	}
	return handle.resp, nil
}

type ZRevRankHandle struct {
	req  *pb.ZRevRankReq
	resp *pb.ZRevRankResp
}

func (h *ZRevRankHandle) Validate() error { return nil }
func (h *ZRevRankHandle) Execute() error {

	resp := &pb.ZRevRankResp{}
	rank, has := db.Db.ZRevRank(h.req.Key, h.req.Member)
	if has {
		resp.Rank = &rank
	}
	h.resp = resp
	log.Log.Debugf("serverHandle:ZRevRank:resp:%v", h.resp)
	return nil
}
