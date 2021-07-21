package handle

import (
	"context"
	"errors"
	"fmt"
	pb "ranking/proto"
	"time"
)

type ZRevRankHandle struct {
	Req  *pb.ZRevRankReq
	Resp *pb.ZRevRankResp
}

func ZRevRank() *ZRevRankHandle { return &ZRevRankHandle{} }
func (h *ZRevRankHandle) Parse(args []string) error {
	args = args[1:]
	if len(args) != 2 {
		return errors.New("args err")
	}
	key := args[0]
	if key == "" {
		return errors.New("err key")
	}
	member := args[1]
	if key == "" {
		return errors.New("err member")
	}

	req := &pb.ZRevRankReq{
		Key: key,
		Member: member,
	}
	h.Req = req
	return nil
}
func (h *ZRevRankHandle) Execute() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := Handle.RankClient.ZRevRank(ctx, h.Req)
	if err != nil {
		return err
	}
	h.Resp = resp
	return nil
}
func (h *ZRevRankHandle) Print() error {
	if h.Resp.Rank == nil {
		fmt.Println("nil")
	}else{
		fmt.Println(h.Resp.String())
	}
	return nil
}
