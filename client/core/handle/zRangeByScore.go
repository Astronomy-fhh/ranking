package handle

import (
	"context"
	"errors"
	"fmt"
	pb "ranking/proto"
	"strconv"
	"time"
)

type ZRangeByScoreHandle struct {
	Req  *pb.ZRangeByScoreReq
	Resp *pb.ZRangeByScoreResp
}

func ZRangeByScore() *ZRangeByScoreHandle { return &ZRangeByScoreHandle{} }
func (h *ZRangeByScoreHandle) Parse(args []string) error {
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

	req := &pb.ZRangeByScoreReq{
		Key: key,
		Start: int64(start),
		End: int64(end),
	}
	h.Req = req
	return nil
}
func (h *ZRangeByScoreHandle) Execute() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := Handle.RankClient.ZRangeByScore(ctx, h.Req)
	if err != nil {
		return err
	}
	h.Resp = resp
	return nil
}
func (h *ZRangeByScoreHandle) Print() error {
	objs := h.Resp.Objs
	if len(objs) > 0 {
		for i, obj := range objs {
			fmt.Printf("%d. %s: %d\n",i,obj.Member,obj.Score)
		}
	}else{
		fmt.Println("nil")
	}
	return nil
}
