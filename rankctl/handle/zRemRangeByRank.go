package handle

import (
	"context"
	"errors"
	"fmt"
	pb "ranking/proto"
	"strconv"
	"time"
)

type ZRemRangeByRankHandle struct {
	Req  *pb.ZRemRangeByRankReq
	Resp *pb.ZRemRangeByRankResp
}

func ZRemRangeByRank() *ZRemRangeByRankHandle { return &ZRemRangeByRankHandle{} }
func (h *ZRemRangeByRankHandle) Parse(args []string) error {
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

	req := &pb.ZRemRangeByRankReq{
		Key: key,
		Start: int64(start),
		End: int64(end),
	}
	h.Req = req
	return nil
}

func (h *ZRemRangeByRankHandle) Execute() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := Handle.RankClient.ZRemRangeByRank(ctx, h.Req)
	if err != nil {
		return err
	}
	h.Resp = resp
	return nil
}

func (h *ZRemRangeByRankHandle) Print() error {
	fmt.Println(h.Resp.String())
	return nil
}
