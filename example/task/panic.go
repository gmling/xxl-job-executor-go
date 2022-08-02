package task

import (
	"context"

	"github.com/gmling/xxl-job-executor-go"
)

func Panic(cxt context.Context, param *xxl.RunReq) (msg string, err error) {
	panic("test panic")
	return
}
