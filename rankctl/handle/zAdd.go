package handle

import (
	"context"
	"errors"
	"fmt"
	pb "ranking/proto"
	"strconv"
	"time"
)

type ZAddHandle struct{
	Req *pb.ZAddReq
	Resp *pb.ZAddResp
}

func ZAdd() *ZAddHandle { return &ZAddHandle{} }
func (h *ZAddHandle) Parse(args []string) error {
	args = args[1:]
	if len(args) < 1 {
		return errors.New("err:syntax error")
	}
	key := args[0]
	if key == "" {
		return errors.New("err:syntax error")
	}
	if len(args)%2 != 1 {
		return errors.New("err:syntax error")
	}
	sIdx := 1
	mIdx := sIdx + 1
	vars := make(map[string]int64, 1)
	for len(args) > mIdx {
		score, err := strconv.ParseInt(args[sIdx], 10, 64)
		if err != nil {
			return errors.New(fmt.Sprintf("parse int err for:%s", args[sIdx]))
		}
		member := args[mIdx]
		vars[member] = score
		sIdx += 2
		mIdx += 2
	}
	req := &pb.ZAddReq{Key: key, Vars: vars}
	h.Req = req
	return nil
}
func (h *ZAddHandle) Execute() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := Handle.RankClient.ZAdd(ctx, h.Req)
	if err != nil {
		return err
	}
	h.Resp = resp
	return nil
}
func (h *ZAddHandle) Print() error {
	fmt.Printf("add:%d,update:%d\n",h.Resp.AddC.Value,h.Resp.UpdateC.Value)
	return nil
}
