package main

import (
	"fmt"
	"time"
)

type atomicInt int

/**
 * 对atomicInt进行加法
 */
func (a *atomicInt) increment() {
	*a++
}

/**
 * 获取a的值
 **/
func (a *atomicInt) get() int {
	return int(*a) // 将atomicInt转换成int再返回
}

func main() {
	var a atomicInt
	a.increment()

	go func() {
		a.increment()
	}()
	time.Sleep(time.Second)
	fmt.Println(a.get()) // 输出2 ，使用go run --race 可以看出有数据冲突，这个get读和increment可能会并发
}
