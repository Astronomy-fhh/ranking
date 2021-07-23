package handle

import (
	"context"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"ranking/db"
	"ranking/log"
	pb "ranking/proto"
)

func (ServerHandle) ZRemRangeByRank(ctx context.Context, req *pb.ZRemRangeByRankReq) (*pb.ZRemRangeByRankResp, error) {
	log.Log.Debugf("serverHandle:ZRemRangeByRank:req:%v", req)
	handle := &ZRemRangeByRankHandle{req: req}
	err := serverHandle.Execute(ctx, handle)
	if err != nil {
		return nil, err
	}
	return handle.resp, nil
}

type ZRemRangeByRankHandle struct {
	req  *pb.ZRemRangeByRankReq
	resp *pb.ZRemRangeByRankResp
}

func (h *ZRemRangeByRankHandle) Validate() error { return nil }
func (h *ZRemRangeByRankHandle) Execute() error {

	ret := db.Db.ZRemRangeByRank(h.req.Key, h.req.Start, h.req.End)
	h.resp = &pb.ZRemRangeByRankResp{
		Ret: wrapperspb.Int64(ret),
	}
	log.Log.Debugf("serverHandle:ZRemRangeByRank:resp:%v", h.resp)
	return nil
}
