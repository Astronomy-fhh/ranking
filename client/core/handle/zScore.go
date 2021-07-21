package handle

import (
	"context"
	"errors"
	"fmt"
	pb "ranking/proto"
	"time"
)

type ZScoreHandle struct {
	Req  *pb.ZScoreReq
	Resp *pb.ZScoreResp
}

func ZScore() *ZScoreHandle { return &ZScoreHandle{} }
func (h *ZScoreHandle) Parse(args []string) error {
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
	req := &pb.ZScoreReq{
		Key: key,
		Member: member,
	}
	h.Req = req
	return nil
}
func (h *ZScoreHandle) Execute() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := Handle.RankClient.ZScore(ctx, h.Req)
	if err != nil {
		return err
	}
	h.Resp = resp
	return nil
}
func (h *ZScoreHandle) Print() error {
	if h.Resp.Score == nil {
		fmt.Println("nil")
	}else{
		fmt.Println(h.Resp.String())
	}
	return nil
}
