package handle

import (
	"context"
	"errors"
	"fmt"
	pb "ranking/proto"
	"strconv"
	"time"
)

type ZAddHandle struct {
	Req  *pb.ZAddReq
	Resp *pb.ZAddResp
}


func ZaddHandle()*ZAddHandle  {
	return &ZAddHandle{}
}

func (h *ZAddHandle) Parse(args []string) error {
	// first args is CMD
	args = args[1:]
	if len(args) < 1 {
		return errors.New("args err")
	}
	key := args[0]
	if key == "" {
		return errors.New("err key")
	}

	if len(args) % 2 != 1 {
		return errors.New("number of mismatches")
	}
	sIdx := 1
	mIdx := sIdx + 1
	val := make(map[string]uint64,1)
	for len(args) > mIdx {
		score,err := strconv.ParseInt(args[sIdx],10,64)
		if err != nil {
			return errors.New(fmt.Sprintf("parse int err for:%s",args[sIdx]))
		}
		member := args[mIdx]
		val[member] = uint64(score)
		sIdx += 2
		mIdx += 2
	}
	req := &pb.ZAddReq{
		Key: key,
		Val: val,
	}
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
	fmt.Println(h.Resp.String())
	return nil
}