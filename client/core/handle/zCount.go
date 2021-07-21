package handle

import (
	"context"
	"errors"
	"fmt"
	pb "ranking/proto"
	"strconv"
	"time"
)

type ZCountHandle struct {
	Req  *pb.ZCountReq
	Resp *pb.ZCountResp
}

func ZCount() *ZCountHandle { return &ZCountHandle{} }
func (h *ZCountHandle) Parse(args []string) error {
	args = args[1:]
	if len(args) != 3 {
		return errors.New("err:syntax error")
	}
	key := args[0]
	if key == "" {
		return errors.New("err:syntax error")
	}
	start, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	end, err := strconv.Atoi(args[2])
	if err != nil {
		return err
	}

	req := &pb.ZCountReq{
		Key: key,
		Start: int64(start),
		End: int64(end),
	}
	h.Req = req
	return nil
}
func (h *ZCountHandle) Execute() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := Handle.RankClient.ZCount(ctx, h.Req)
	if err != nil {
		return err
	}
	h.Resp = resp
	return nil
}

func (h *ZCountHandle) Print() error {
	if h.Resp.Ret == nil {
		fmt.Println("nil")
	}else{
		fmt.Println(h.Resp.String())
	}
	return nil
}
