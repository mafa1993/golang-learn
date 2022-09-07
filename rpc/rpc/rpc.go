package rpcc

import (
	"errors"
	"fmt"
)

// 调用方式 service.Method
type DemoService struct {
}

type Args struct {
	A, B int
}

/**
 * 创建服务，服务必须有两个参数，一个是参数，一个是结果, 结果必须为指针类型
 */
func (DemoService DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("0不能做除数")
	}

	*result = float64(args.A) / float64(args.B)
	fmt.Printf("%d,%d,%f", args.A, args.B, *result)

	return nil
}
