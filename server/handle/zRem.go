package handle

import (
	"context"
	"ranking/db"
	"ranking/log"
	pb "ranking/proto"
)

func (ServerHandle) ZRem(ctx context.Context, req *pb.ZRemReq) (*pb.ZRemResp, error) {
	log.Log.Debugf("serverHandle:ZRem:req:%v", req)
	handle := &ZRemHandle{req: req}
	err := serverHandle.Execute(ctx, handle)
	if err != nil {
		return nil, err
	}
	return handle.resp, nil
}

type ZRemHandle struct {
	req *pb.ZRemReq
	resp *pb.ZRemResp
}

func (h *ZRemHandle) Validate() error { return nil }
func (h *ZRemHandle) Execute() error {

	ret := db.Db.ZRem(h.req.Key, h.req.Members)
	h.resp = &pb.ZRemResp{
		Ret: &ret,
	}
	log.Log.Debugf("serverHandle:ZRem:resp:%v", h.resp)
	return nil
}
