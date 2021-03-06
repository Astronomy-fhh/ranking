package client

import (
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"ranking/client/core/handle"
	"ranking/config"
	"ranking/log"
	pb "ranking/proto"
	"syscall"
)

type Client struct {
   Cmd *Cmd
   Handle *handle.ClientHandle
   Config *config.ClientConfig
}

func NewClient(config *config.ClientConfig)*Client {
	client := &Client{}
	client.Cmd = NewCmdApp()
	client.InitConfig(config)
	client.InitHandle()
	return client
}


func (c *Client) InitConfig(config *config.ClientConfig)  {
	c.Config = config
}


func (c *Client) InitHandle()  {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	grpcServerAddr := c.Config.HttpAddr
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(grpcServerAddr, opts...)
	if err != nil {
		log.Log.Fatalf("fail to dial:%v",err.Error())
	}
	RankClient := pb.NewRankClient(conn)
	clientHandle := &handle.ClientHandle{
		ClientConn: conn,
		RankClient: RankClient,
	}
	handle.Handle = clientHandle
	c.Handle = clientHandle
}


func (c *Client) Run()  {
	go c.RunCmd()

	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	select {
	case s := <-chSignal:
		log.Log.Infof("signal received: %v", s)
		signal.Reset(syscall.SIGINT, syscall.SIGTERM)
	}

	c.shutdown()
}

func (c *Client) shutdown()  {
	_ = c.Handle.ClientConn.Close()
}


func (c *Client) RunCmd()  {
	c.Cmd.RunCmd(c.Handle.ClientConn.Target())
}



