package command

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"os"
	"ranking/rankctl/handle"
	"regexp"
	"strings"
)

type Cmd struct {
	CmdApp *cli.App
	CmdAppArgs []string
}

func NewCmdApp()*Cmd {
	CmdApp := &cli.App{
		Name:    "rank-cli",
		Version: "1.0.0",
		EnableBashCompletion: true,
		UsageText: "command [command options] [arguments...]",
	}
	cmd := &Cmd{}
	cmd.CmdApp = CmdApp
	cmd.InitCommends()
	return cmd
}



func (c *Cmd) InitCommends()  {
	c.CmdApp.Commands = []*cli.Command{
		{
			Name: "exit",
			Aliases: []string{"q"},
			Usage: "exit",
			Action: func(context *cli.Context) error {
				return cli.Exit("",0)
			},
		},
		{
			Name: "zadd",
			Aliases: []string{"ZADD"},
			Usage: "ZADD key score member [score member ...]",
			Action: func(context *cli.Context) error {
				err := handle.CmdHandle(handle.ZAdd(),c.CmdAppArgs)
				return err
			},
		},
		{
			Name: "zcard",
			Aliases: []string{"ZCARD"},
			Usage: "ZCARD key ",
			Action: func(context *cli.Context) error {
				err := handle.CmdHandle(handle.ZCard(),c.CmdAppArgs)
				return err
			},
		},
		{
			Name: "zcount",
			Aliases: []string{"ZCOUNT"},
			Usage: "ZCOUNT key min max ",
			Action: func(context *cli.Context) error {
				err := handle.CmdHandle(handle.ZCount(),c.CmdAppArgs)
				return err
			},
		},
		{
			Name: "zincrby",
			Aliases: []string{"ZINCRBY"},
			Usage: "ZINCRBY key incr member ",
			Action: func(context *cli.Context) error {
				err := handle.CmdHandle(handle.ZIncrBy(),c.CmdAppArgs)
				return err
			},
		},
		{
			Name: "zrange",
			Aliases: []string{"ZRANGE"},
			Usage: "zrange key start stop ",
			Action: func(context *cli.Context) error {
				err := handle.CmdHandle(handle.ZRange(),c.CmdAppArgs)
				return err
			},
		},
		{
			Name: "zrangebyscore",
			Aliases: []string{"ZRANGEBYSCORE"},
			Usage: "zrangebyscore key start stop ",
			Action: func(context *cli.Context) error {
				err := handle.CmdHandle(handle.ZRangeByScore(),c.CmdAppArgs)
				return err
			},
		},
		{
			Name: "zrank",
			Aliases: []string{"ZRANK"},
			Usage: "zrank key member ",
			Action: func(context *cli.Context) error {
				err := handle.CmdHandle(handle.ZRank(),c.CmdAppArgs)
				return err
			},
		},
		{
			Name: "zrem",
			Aliases: []string{"ZREM"},
			Usage: "zrem key member [member ...] ",
			Action: func(context *cli.Context) error {
				err := handle.CmdHandle(handle.ZRem(),c.CmdAppArgs)
				return err
			},
		},
		{
			Name: "zremrangebyrank",
			Aliases: []string{"ZREMRANGEBYRANK"},
			Usage: "zremrangebyrank key start stop ",
			Action: func(context *cli.Context) error {
				err := handle.CmdHandle(handle.ZRemRangeByRank(),c.CmdAppArgs)
				return err
			},
		},
		{
			Name: "zremrangebyscore",
			Aliases: []string{"ZREMRANGEBYSCORE"},
			Usage: "zremrangebyscore key start stop ",
			Action: func(context *cli.Context) error {
				err := handle.CmdHandle(handle.ZRemRangeByScore(),c.CmdAppArgs)
				return err
			},
		},
		{
			Name: "zrevrange",
			Aliases: []string{"ZRANGE"},
			Usage: "zrevrange key start stop ",
			Action: func(context *cli.Context) error {
				err := handle.CmdHandle(handle.ZRevRange(),c.CmdAppArgs)
				return err
			},
		},
		{
			Name: "zrevrangebyscore",
			Aliases: []string{"ZRANGEBYSCORE"},
			Usage: "zrevrangebyscore key start stop ",
			Action: func(context *cli.Context) error {
				err := handle.CmdHandle(handle.ZRevRangeByScore(),c.CmdAppArgs)
				return err
			},
		},
		{
			Name: "zrevrank",
			Aliases: []string{"ZREVRANK"},
			Usage: "zrevrank key start stop ",
			Action: func(context *cli.Context) error {
				err := handle.CmdHandle(handle.ZRevRank(),c.CmdAppArgs)
				return err
			},
		},
		{
			Name: "zscore",
			Aliases: []string{"ZSCORE"},
			Usage: "zscore key member ",
			Action: func(context *cli.Context) error {
				err := handle.CmdHandle(handle.ZScore(),c.CmdAppArgs)
				return err
			},
		},
	}
}

func (c *Cmd) RunCmd(connInfo string)  {

	reader := bufio.NewReader(os.Stdin)
	placeInfo := connInfo + "> "

	for  {
		GreenPrint(placeInfo)
		textLine, _ := reader.ReadString('\n')
		textLine = strings.TrimSpace(textLine)

		compile := regexp.MustCompile("\\s+")
		args := compile.Split(textLine, -1)
		if  args[0] == "" {
			continue
		}

		command := c.CmdApp.Command(args[0])
		if command == nil {
			RedPrint("err: unknown command [" + args[0] + "]\n")
			continue
		}
		c.CmdAppArgs = args
		argsSlice := make([]string,1)
		argsSlice = append(argsSlice,args[0:]...)
		err := c.CmdApp.Run(argsSlice)
		if err != nil {
			RedPrint(fmt.Sprintf("err:%s\n",err.Error()))
		}
	}
}

func GreenPrint(str string)  {
	c := color.New(color.FgGreen,color.Bold).FprintfFunc()
	c(os.Stdout, str)
}
func RedPrint(str string)  {
	c := color.New(color.FgRed).FprintfFunc()
	c(os.Stdout, str)
}
