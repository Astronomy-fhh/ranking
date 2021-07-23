package handle

import (
	"context"
	"errors"
	"fmt"
	pb "ranking/proto"
	"time"
)

type ZRankHandle struct {
	Req  *pb.ZRankReq
	Resp *pb.ZRankResp
}

func ZRank() *ZRankHandle { return &ZRankHandle{} }
func (h *ZRankHandle) Parse(args []string) error {
	args = args[1:]
	if len(args) != 2 {
		return errors.New("err:syntax error")
	}
	key := args[0]
	if key == "" {
		return errors.New("err:syntax error")
	}
	member := args[1]
	if key == "" {
		return errors.New("err:syntax error")
	}

	req := &pb.ZRankReq{
		Key: key,
		Member: member,
	}
	h.Req = req
	return nil
}
func (h *ZRankHandle) Execute() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := Handle.RankClient.ZRank(ctx, h.Req)
	if err != nil {
		return err
	}
	h.Resp = resp
	return nil
}
func (h *ZRankHandle) Print() error {
	if h.Resp.Rank == nil {
		fmt.Println("nil")
	}else{
		fmt.Println(h.Resp.Rank.Value)
	}
	return nil
}
