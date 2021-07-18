package core

import (
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"ranking/config"
	"ranking/db"
	"ranking/log"
	pb "ranking/proto"
	"syscall"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewServer()*Server {
	rs := &Server{}
	return rs
}


func (rs *Server) Run()  {

	db.InitDB()
	InitStatusServer()

	lis, err := net.Listen("tcp", config.SConfig.HttpAddr)
	if err != nil {
		log.Log.Fatalf("net listen for addr:%s, error:%s",config.SConfig.HttpAddr,err.Error())
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

func (rs *Server) Shutdown()  {

}







