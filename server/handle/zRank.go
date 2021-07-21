package handle

import (
	"context"
	"ranking/db"
	"ranking/log"
	pb "ranking/proto"
)

func (ServerHandle) ZRank(ctx context.Context, req *pb.ZRankReq) (*pb.ZRankResp, error) {
	log.Log.Debugf("serverHandle:ZRank:req:%v", req)
	handle := &ZRankHandle{req: req}
	err := serverHandle.Execute(ctx, handle)
	if err != nil {
		return nil, err
	}
	return handle.resp, nil
}

type ZRankHandle struct {
	req  *pb.ZRankReq
	resp *pb.ZRankResp
}

func (h *ZRankHandle) Validate() error { return nil }
func (h *ZRankHandle) Execute() error {

	rank, has := db.Db.ZRank(h.req.Key, h.req.Member)
	resp := &pb.ZRankResp{}
	if has {
		resp.Rank = &rank
	}
	h.resp = resp
	log.Log.Debugf("serverHandle:ZRank:resp:%v", h.resp)
	return nil
}
