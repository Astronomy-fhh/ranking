package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"ranking/config"
	"ranking/container"
	"ranking/log"
	pb "ranking/proto"
	"sync"
	"syscall"
)

type RankServer struct {
	containers map[string]*container.Container
	config     *config.ServerConfig
	grpcServer *grpc.Server
	sync.Mutex
	pb.UnimplementedRankServer
}

func NewRankServer()*RankServer  {
	rs := &RankServer{
		containers: make(map[string]*container.Container),
	}
	return rs
}

func (rs *RankServer) SetConfig(config *config.ServerConfig)  {
	rs.config = config
}

func (rs *RankServer) Run()  {
	lis, err := net.Listen("tcp", rs.config.HttpAddr)
	if err != nil {
		log.Log.Fatalf("net listen for addr:%s, error:%s",rs.config.HttpAddr,err.Error())
	}
	s := grpc.NewServer()
	rs.grpcServer = s
	pb.RegisterRankServer(s, rs)
	if err := s.Serve(lis); err != nil {
		log.Log.Fatalf("listener lis fail, error:%s",err.Error())
	}

	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-chSignal:
		log.Log.Infof("signal received: %v", s)
		signal.Reset(syscall.SIGINT, syscall.SIGTERM)
	}

}

func (rs *RankServer) Shutdown()  {
	rs.grpcServer.Stop()
}

func (rs *RankServer) getOrInitContainer(key string)*container.Container {
	myContainer := rs.containers[key]
	if myContainer == nil{
		rs.Lock()
		defer rs.Unlock()
		set := rs.containers[key]
		if set == nil{
			rs.containers[key] = container.NewContainer()
		}
	}
	return rs.containers[key]
}

func (rs *RankServer) ZAdd(ctx context.Context,req *pb.ZAddReq) (*pb.ZAddResp, error)  {
	return &pb.ZAddResp{Ret: uint64(len(req.Val))},nil
}




