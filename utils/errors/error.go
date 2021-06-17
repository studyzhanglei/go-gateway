package errors

import (
	"github.com/studyzhanglei/grpc-proto/pb/common"
	"google.golang.org/grpc/codes"
	"reflect"
	"runtime"
)


const (
	OK      = 0
	SUCCESS = "success"
)

type errorCode interface {
	String() string
}

type Error struct {
	code int64
	err  string
	*stack
}

func (e *Error) Error() string { return e.err }

func (e *Error) String() string { return e.err }

func NewFromCode(code interface{}) *Error {
	return &Error{
		code:  reflect.ValueOf(code).Int(),
		err:   code.(errorCode).String(),
		stack: callers(),
	}
}


// stack represents a stack of program counters.
type stack []uintptr


func (e *Error) ResHeader() *common.BusinessStatus {
	return &common.BusinessStatus{
		Code: int32(e.code),
		Msg:     e.err,
	}
}

func (e *Error) GetCode() int64 {
	return e.code
}

func GetResHeader(err error) *common.BusinessStatus {
	if err == nil {
		return &common.BusinessStatus{
			Code: OK,
			Msg: SUCCESS,
		}
	}
	e, ok := err.(*Error)
	if ok {
		return e.ResHeader()
	}

	return &common.BusinessStatus{
		Code: int32(codes.Unknown),
		Msg:     err.Error(),
	}
}

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}