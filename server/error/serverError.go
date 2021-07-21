package serverError //server err

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func NewError(code codes.Code,msg string,args... interface{})error{
	if len(args) > 0 {
		msg = fmt.Sprintf(msg,args...)
	}
	s := status.New(code, msg)
	return s.Err()
}

func NewInvalidArgumentError(msg string)error  {
	return NewError(codes.InvalidArgument,msg)
}