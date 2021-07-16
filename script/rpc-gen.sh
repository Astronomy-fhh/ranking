#!/usr/bin/env bash

# You must specify a --proto_path which encompasses this file
cd ..
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/rpc.proto
cd ./script