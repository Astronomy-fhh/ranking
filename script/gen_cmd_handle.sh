#!/usr/bin/env bash

if [[ $# != 1 ]]; then
    echo "Usage: ./`basename $0` handler-name"
    exit 1
fi

handleName=$1
upFirstName=`echo "$1" | perl -F'\n' -lane 'print ucfirst($F[0])'`

upFirstName="Z${upFirstName}"

fileName="../client/core/handle/z${handleName}.go"

if [[ -f "$fileName" ]]; then
    echo "$REQ_FILE_FROM_ROOT already exists."
    exit 1
fi

read -d '' fileContent  << EOF
package handle

import (
	"context"
	"errors"
	"fmt"
	pb "ranking/proto"
	"strconv"
	"time"
)

type ${upFirstName}Handle struct {
	Req  *pb.${upFirstName}Req
	Resp *pb.${upFirstName}Resp
}

func ${upFirstName}() *${upFirstName}Handle {
	return &${upFirstName}Handle{}
}

func (h *${upFirstName}Handle) Parse(args []string) error {
	args = args[1:]
	if len(args) < 1 {
		return errors.New("args err")
	}
	key := args[0]
	if key == "" {
		return errors.New("err key")
	}


	req := &pb.${upFirstName}Req{

	}
	h.Req = req
	return nil
}

func (h *${upFirstName}Handle) Execute() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := Handle.RankClient.${upFirstName}(ctx, h.Req)
	if err != nil {
		return err
	}
	h.Resp = resp
	return nil
}

func (h *${upFirstName}Handle) Print() error {
	fmt.Println(h.Resp.String())
	return nil
}

EOF

echo $fileContent > "$fileName"
echo "$fileName generated."
