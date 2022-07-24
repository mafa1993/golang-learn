package main

import (
	"fmt"
	"sync"
	"time"
)

// 解决 atomic.go 冲突的问题，使用 sync.Mutex

type atomicInt struct {
	value int
	lock  sync.Mutex // 锁
}

/**
 * 对atomicInt进行加法
 */
func (a *atomicInt) increment() {
	a.lock.Lock()         // 加锁
	defer a.lock.Unlock() // 解锁
	a.value++

}

/**
 * 获取a的值
 **/
func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return int(a.value) // 将atomicInt转换成int再返回
}

func main() {
	var a atomicInt
	a.increment()

	go func() {
		a.increment()
	}()
	time.Sleep(time.Second)
	fmt.Println(a.get()) 
}
