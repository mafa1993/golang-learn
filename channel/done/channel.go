package main

import (
	"fmt"
	"time"
)

// 使用chan 来通知外面协程结束
type worker struct {
	in   chan int
	done chan bool
}

func chanDemo() {

	var workers [10]worker
	for i := 0; i < 10; i++ {
		// 创建10个worker，没有向chan里传递数据，不会执行
		workers[i] = createWorker(i)

	}

	for i := 0; i < 10; i++ {
		// 传递数据，使得worker执行
		workers[i].in <- i
		// <-workers[i].done 如果在这里接收，会阻塞等待里面的done发出数据，会导致10个worker的执行不是并发的，顺序执行
	}

	for _, v := range workers {
		<-v.done
	}

}

func main() {
	chanDemo()
	time.Sleep(time.Second) // 防止还没运行，就退出了
}

// 增加done chan 用来通知外面，已经处理完毕
func doWork(id int, c chan int, done chan bool) {
	for {
		rlt, ok := <-c // 用两个值进行收

		// 如果没有收到数据，就进入下次循环
		if !ok {
			break
		}
		fmt.Printf("id%d接收到的是%d\n", id, rlt)
		// done <- true 这里如果直接发，外面没有接收的，会死锁

		// 利用闭包，暂存起来，当外面有人从cone里接收值的时候，触发
		go func() {
			done <- true
		}()
	}

}

// 创建worker，利用闭包，每个协程固定数据
func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}

	go doWork(id, w.in, w.done)

	return w
}
