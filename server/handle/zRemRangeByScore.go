package handle

import (
	"context"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"ranking/db"
	"ranking/log"
	pb "ranking/proto"
)

func (ServerHandle) ZRemRangeByScore(ctx context.Context, req *pb.ZRemRangeByScoreReq) (*pb.ZRemRangeByScoreResp, error) {
	log.Log.Debugf("serverHandle:ZRemRangeByScore:req:%v", req)
	handle := &ZRemRangeByScoreHandle{req: req}
	err := serverHandle.Execute(ctx, handle)
	if err != nil {
		return nil, err
	}
	return handle.resp, nil
}

type ZRemRangeByScoreHandle struct {
	req  *pb.ZRemRangeByScoreReq
	resp *pb.ZRemRangeByScoreResp
}

func (h *ZRemRangeByScoreHandle) Validate() error { return nil }
func (h *ZRemRangeByScoreHandle) Execute() error {

	ret := db.Db.ZRemRangeByScore(h.req.Key, h.req.Start, h.req.End)
	h.resp = &pb.ZRemRangeByScoreResp{
		Ret: wrapperspb.Int64(ret),
	}
	log.Log.Debugf("serverHandle:ZRemRangeByScore:resp:%v", h.resp)
	return nil
}
