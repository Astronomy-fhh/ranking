package handle

import (
	"context"
	"errors"
	"fmt"
	pb "ranking/proto"
	"strconv"
	"time"
)

type ZRevRangeByScoreHandle struct {
	Req  *pb.ZRevRangeByScoreReq
	Resp *pb.ZRevRangeByScoreResp
}

func ZRevRangeByScore() *ZRevRangeByScoreHandle { return &ZRevRangeByScoreHandle{} }
func (h *ZRevRangeByScoreHandle) Parse(args []string) error {
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

	req := &pb.ZRevRangeByScoreReq{
		Key: key,
		Start: int64(start),
		End: int64(end),
	}
	h.Req = req
	return nil
}
func (h *ZRevRangeByScoreHandle) Execute() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := Handle.RankClient.ZRevRangeByScore(ctx, h.Req)
	if err != nil {
		return err
	}
	h.Resp = resp
	return nil
}
func (h *ZRevRangeByScoreHandle) Print() error {
	objs := h.Resp.Objs
	if len(objs) > 0 {
		for i, obj := range objs {
			fmt.Printf("%d. %s: %d\n",i,obj.Member,obj.Score)
		}
	}else{
		fmt.Println("empty!")
	}
	return nil
}
