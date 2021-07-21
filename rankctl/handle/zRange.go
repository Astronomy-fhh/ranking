package handle

import (
	"context"
	"errors"
	"fmt"
	pb "ranking/proto"
	"strconv"
	"time"
)

type ZRangeHandle struct {
	Req  *pb.ZRangeReq
	Resp *pb.ZRangeResp
}

func ZRange() *ZRangeHandle { return &ZRangeHandle{} }
func (h *ZRangeHandle) Parse(args []string) error {
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

	req := &pb.ZRangeReq{
		Key: key,
		Start: int64(start),
		End: int64(end),
	}
	h.Req = req
	return nil
}
func (h *ZRangeHandle) Execute() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := Handle.RankClient.ZRange(ctx, h.Req)
	if err != nil {
		return err
	}
	h.Resp = resp
	return nil
}
func (h *ZRangeHandle) Print() error {
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
