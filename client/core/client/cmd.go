package client

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"os"
	"ranking/client/core/handle"
	pb "ranking/proto"
	"regexp"
	"strings"
)

type Cmd struct {
	CmdApp *cli.App
	CmdAppArgs []string
}

func NewCmdApp()*Cmd  {
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
			Name: "ZADD",
			Aliases: []string{"ZADD"},
			Usage: "ZADD key score member [score member ...]",
			Action: func(context *cli.Context) error {
				err := handle.HandOut(pb.Method_ZADD, c.CmdAppArgs)
				return err
			},
		},
		{
			Name: "qqq",
			Aliases: []string{"rc"},
			Usage: "rank-cli",
			Action: func(context *cli.Context) error {
				fmt.Println("aaaa!!!")
				return nil
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
			RedPrint("ERR unknown command\n")
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
