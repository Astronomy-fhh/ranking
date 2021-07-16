package server

import (
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"ranking/config"
	"ranking/log"
	pb "ranking/proto"
	"syscall"
)

type RankServer struct {
	config     *config.ServerConfig
	grpcServer *grpc.Server
}

func NewRankServer()*RankServer {
	rs := &RankServer{}
	return rs
}

func (rs *RankServer) SetConfig(config *config.ServerConfig)  {
	rs.config = config
}

func (rs *RankServer) Run()  {
	ServerDataInit()
	lis, err := net.Listen("tcp", rs.config.HttpAddr)
	if err != nil {
		log.Log.Fatalf("net listen for addr:%s, error:%s",rs.config.HttpAddr,err.Error())
	}
	s := grpc.NewServer()
	rs.grpcServer = s
	pb.RegisterRankServer(s, NewServerHandle())

	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		select {
		case s := <-chSignal:
			log.Log.Infof("signal received: %v", s)
			signal.Reset(syscall.SIGINT, syscall.SIGTERM)
			os.Exit(0)
		}
	}()

	if err := s.Serve(lis); err != nil {
		log.Log.Fatalf("listener lis fail, error:%s",err.Error())
	}
}

func (rs *RankServer) Shutdown()  {

}







