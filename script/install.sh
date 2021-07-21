#!/usr/bin/env bash


CUR_DIR=$(cd `dirname $0` ; pwd)
cd $CUR_DIR
cd ../server
go build -o ../bin/
cd ../rankctl
go build -o ../bin/
echo "ok"