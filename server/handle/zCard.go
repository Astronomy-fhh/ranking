package handle

import (
	"context"
	"ranking/db"
	"ranking/log"
	pb "ranking/proto"
)

func (ServerHandle)ZCard(ctx context.Context,req *pb.ZCardReq)(*pb.ZCardResp,error) {
	log.Log.Debugf("serverHandle:ZAdd:req:%v",req)

	handle := &ZCardHandle{req: req}
	err := serverHandle.Execute(ctx, handle)
	if err != nil {
		return nil, err
    }
	return handle.resp,nil
}

type ZCardHandle struct {
	req *pb.ZCardReq
	resp *pb.ZCardResp
}

func (h *ZCardHandle) Validate()error  {
	return nil
}

func (h *ZCardHandle) Execute()error  {
	resp := &pb.ZCardResp{}
	card, has := db.Db.ZCard(h.req.Key)
	if has {
		resp.Ret = &card
	}
	h.resp = resp
	log.Log.Debugf("serverHandle:ZCard:resp:%v",resp)
	return nil
}

