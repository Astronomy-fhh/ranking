package e

import (
	"fmt"
	pb "ranking/proto"
)

type Error struct {
	Code pb.ErrorNo
	Msg string
}

func New(code pb.ErrorNo,msg string,args ...interface{})*Error  {
	if len(args) != 0 {
		msg= fmt.Sprintf(msg,args...)
	}
     return &Error{
     	Code: code,
     	Msg: msg,
	 }
}


func (e *Error) Error()string  {
	return e.Msg
}