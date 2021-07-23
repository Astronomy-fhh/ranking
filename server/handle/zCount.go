package handle

import (
	"context"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"ranking/db"
	"ranking/log"
	pb "ranking/proto"
)

func (ServerHandle) ZCount(ctx context.Context, req *pb.ZCountReq) (*pb.ZCountResp, error) {
	log.Log.Debugf("serverHandle:ZCount:req:%v", req)
	handle := &ZCountHandle{req: req}
	err := serverHandle.Execute(ctx, handle)
	if err != nil {
		return nil, err
	}
	return handle.resp, nil
}

type ZCountHandle struct {
	req  *pb.ZCountReq
	resp *pb.ZCountResp
}

func (h *ZCountHandle) Validate() error { return nil }
func (h *ZCountHandle) Execute() error {
	resp := &pb.ZCountResp{
	}
	count, has := db.Db.ZCount(h.req.Key, h.req.Start, h.req.End)
	if has {
		resp.Ret = wrapperspb.Int64(count)
	}
	h.resp = resp
	log.Log.Debugf("serverHandle:ZCount:resp:%v", h.resp)
	return nil
}
