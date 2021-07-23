package handle

import (
	"context"
	"errors"
	"fmt"
	pb "ranking/proto"
	"strconv"
	"time"
)

type ZIncrByHandle struct {
	Req  *pb.ZIncrByReq
	Resp *pb.ZIncrByResp
}

func ZIncrBy() *ZIncrByHandle { return &ZIncrByHandle{} }
func (h *ZIncrByHandle) Parse(args []string) error {
	args = args[1:]
	if len(args) != 3 {
		return errors.New("err:syntax error")
	}
	key := args[0]
	if key == "" {
		return errors.New("err:syntax error")
	}
	incr, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	member := args[2]
	if key == "" {
		return errors.New("err:syntax error")
	}
	req := &pb.ZIncrByReq{
		Key:key,
		Incr: int64(incr),
		Member: member,
	}
	h.Req = req
	return nil
}
func (h *ZIncrByHandle) Execute() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := Handle.RankClient.ZIncrBy(ctx, h.Req)
	if err != nil {
		return err
	}
	h.Resp = resp
	return nil
}
func (h *ZIncrByHandle) Print() error {
	fmt.Println(h.Resp.Ret.Value)
	return nil
}
