package main

import (
	"fmt"
	"runtime"
	"time"
)

// 并发编程一
func printHello() {
	fmt.Println("hello")
}

func main() {
	for i := 0; i < 10; i++ {
		go printHello()
	}

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Printf("hello%d\n", i)
		}(i)
	}

	// ========= 无io操作  死循环演示
	// var arr [10]int
	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		arr[i]++
	// 		// runtime.Gosched()
	// 	}(i)
	// }

	// fmt.Println(arr) // 可以看到 协程切换了多少次

	// ========= out of range 演示
	var arr1 [10]int
	for i := 0; i < 10; i++ {
		go func() {
			arr1[i]++   // 这个i和外面的i是同一个i，同时，for进行写，这里读 发生了冲突
			runtime.Gosched()
		}()
	}


	// 如果没有等待，main会先退出，可能什么也打不出来
	time.Sleep(time.Second)

}
