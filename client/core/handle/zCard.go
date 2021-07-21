package handle

import (
	"context"
	"errors"
	"fmt"
	pb "ranking/proto"
	"time"
)

type ZCardHandle struct {
	Req  *pb.ZCardReq
	Resp *pb.ZCardResp
}


func ZCard()*ZCardHandle  {
	return &ZCardHandle{}
}

func (h *ZCardHandle) Parse(args []string) error {
	// first args is CMD
	args = args[1:]
	if len(args) < 1 ||len(args) > 2 {
		return errors.New("err:syntax error")
	}
	key := args[0]
	if key == "" {
		return errors.New("err:syntax error")
	}
	req := &pb.ZCardReq{
		Key: key,
	}
	h.Req = req
	return nil
}

func (h *ZCardHandle) Execute() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := Handle.RankClient.ZCard(ctx, h.Req)
	if err != nil {
		return err
	}
	h.Resp = resp
	return nil
}


func (h *ZCardHandle) Print() error {
	if h.Resp.Ret == nil {
		fmt.Println("nil")
	}else{
		fmt.Println(h.Resp.String())
	}
	return nil
}
