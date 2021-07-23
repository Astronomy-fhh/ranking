package main

import (
	"log"
	"os"

	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
	"github.com/golang/protobuf/proto"
	pb "ranking/proto"
)

func main() {
	// 组装BinaryData
	item := pb.ZIncrByReq{
		Key: "testK",
		Member: "testM",
		Incr: 1,
	}
	buf := proto.Buffer{}
	err := buf.EncodeMessage(&item)
	if err != nil {
		log.Fatal(err)
		return
	}
	report, err := runner.Run(
		"message.Rank.ZIncrBy", //  'package.Service/method' or 'package.Service.Method'
		":11917",
		runner.WithProtoFile("/Users/fanhuhu/PhpstormProjects/GOPATH/src/ranking/proto/rpc.proto", []string{}),
		runner.WithBinaryData(buf.Bytes()),
		runner.WithInsecure(true),
		runner.WithTotalRequests(100000),
		// 并发参数
		runner.WithConcurrencySchedule(runner.ScheduleLine),
		runner.WithConcurrencyStep(100),
		runner.WithConcurrencyStart(5),
		runner.WithConcurrencyEnd(1000),

	)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 指定输出路径
	file, err := os.Create("report.html")
	if err != nil {
		log.Fatal(err)
		return
	}
	rp := printer.ReportPrinter{
		Out:    file,
		Report: report,
	}
	// 指定输出格式
	_ = rp.Print("html")
}