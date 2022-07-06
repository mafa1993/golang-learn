package main

import (
	"errors"
	"fmt"
)

func tryRecover() {
	defer func() {
		rlt := recover()
		//fmt.Printf("%T", rlt)
		if err, ok := rlt.(error); ok {
			fmt.Println("err is ", err)
		} else {
			panic(rlt) // 如果遇到无法处理的错误，重新panic
		}
	}() // 自执行函数

	b := 0
	a := 4 / b
	fmt.Println(a)
	panic(errors.New("aaa"))
}

func main() {
	tryRecover()
}
