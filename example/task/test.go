package task

import (
	"context"
	"fmt"

	"github.com/xxl-job/xxl-job-executor-go"
)

func Test(cxt context.Context, param *xxl.RunReq) (msg string, err error) {
	fmt.Println("test one task" + param.ExecutorHandler + " param：" + param.ExecutorParams + " log_id:" + xxl.Int64ToStr(param.LogID))
	return "test done", nil
}
