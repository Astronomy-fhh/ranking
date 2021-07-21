#!/usr/bin/env bash

if [[ $# != 1 ]]; then
    echo "Usage: ./`basename $0` handler-name"
    exit 1
fi

handleName=$1
upFirstName=`echo "$1" | perl -F'\n' -lane 'print ucfirst($F[0])'`

upFirstName="Z${upFirstName}"

fileName="../server/core/handle/z${handleName}.go"

if [[ -f "$fileName" ]]; then
    echo "$REQ_FILE_FROM_ROOT already exists."
    exit 1
fi

read -d '' fileContent  << EOF
package handle

import (
	"context"
	"ranking/config"
	"ranking/db"
	e "ranking/error"
	"ranking/log"
	pb "ranking/proto"
)

func (ServerHandle)${upFirstName}(ctx context.Context,req *pb.${upFirstName}Req)(*pb.${upFirstName}Resp,error) {
	log.Log.Debugf("serverHandle:${upFirstName}:req:%v",req)

	handle := &${upFirstName}Handle{req: req}
	err := serverHandle.Execute(ctx, handle)
	if err != nil {
		return nil, err
    }
	return handle.resp,nil
}

type ${upFirstName}Handle struct {
	req *pb.${upFirstName}Req
	resp *pb.${upFirstName}Resp
}

func (h *${upFirstName}Handle) Validate()error  {
	return nil
}

func (h *${upFirstName}Handle) Execute()error  {
	log.Log.Debugf("serverHandle:${upFirstName}:ret:%v",h.resp)
	return nil
}
EOF

echo $fileContent > "$fileName"
echo "$fileName generated."
