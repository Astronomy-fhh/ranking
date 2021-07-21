package handle

import (
	"context"
	"errors"
	"fmt"
	pb "ranking/proto"
	"strconv"
	"time"
)

type ZRevRangeHandle struct {
	Req  *pb.ZRevRangeReq
	Resp *pb.ZRevRangeResp
}

func ZRevRange() *ZRevRangeHandle { return &ZRevRangeHandle{} }
func (h *ZRevRangeHandle) Parse(args []string) error {
	args = args[1:]
	if len(args) != 3 {
		return errors.New("args err")
	}
	key := args[0]
	if key == "" {
		return errors.New("err key")
	}
	start, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	end, err := strconv.Atoi(args[2])
	if err != nil {
		return err
	}

	req := &pb.ZRevRangeReq{
		Key: key,
		Start: int64(start),
		End: int64(end),
	}
	h.Req = req
	return nil
}
func (h *ZRevRangeHandle) Execute() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := Handle.RankClient.ZRevRange(ctx, h.Req)
	if err != nil {
		return err
	}
	h.Resp = resp
	return nil
}
func (h *ZRevRangeHandle) Print() error {
	objs := h.Resp.Objs
	if objs != nil {
		for i, obj := range objs {
			fmt.Printf("%d. %s: %d\n",i,obj.Member,obj.Score)
		}
	}else{
		fmt.Println("nil")
	}
	return nil
}