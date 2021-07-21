package handle

import (
	"context"
	"errors"
	"fmt"
	pb "ranking/proto"
	"time"
)

type ZRemHandle struct {
	Req  *pb.ZRemReq
	Resp *pb.ZRemResp
}

func ZRem() *ZRemHandle { return &ZRemHandle{} }
func (h *ZRemHandle) Parse(args []string) error {
	args = args[1:]
	if len(args) < 2 {
		return errors.New("err:syntax error")
	}
	key := args[0]
	if key == "" {
		return errors.New("err:syntax error")
	}

	members := make([]string,0)
	members = append(members,args[1:]...)
	req := &pb.ZRemReq{
		Key: key,
		Members: members,
	}
	h.Req = req
	return nil
}
func (h *ZRemHandle) Execute() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := Handle.RankClient.ZRem(ctx, h.Req)
	if err != nil {
		return err
	}
	h.Resp = resp
	return nil
}
func (h *ZRemHandle) Print() error {
	fmt.Println(h.Resp.String())
	return nil
}
