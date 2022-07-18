package main

import (
	"fmt"
	"sync"
)

// 使用sync.WaitGroup实现等待
type worker struct {
	in   chan int
	wg   *sync.WaitGroup // 内外的共用，所以使用指针类型，不能重新创建一个
	done func()          // 把done封装起来 直接调用
}

func chanDemo() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		// 创建10个worker，没有向chan里传递数据，不会执行
		workers[i] = createWorker(i, &wg)
	}

	for i := 0; i < 10; i++ {
		// 传递数据，使得worker执行
		workers[i].in <- i
		// <-workers[i].done 如果在这里接收，会阻塞等待里面的done发出数据，会导致10个worker的执行不是并发的，顺序执行

		// 每个任务+1
		wg.Add(1)
	}
	for i := 0; i < 10; i++ {
		// 传递数据，使得worker执行
		workers[i].in <- i + 10
		// <-workers[i].done 如果在这里接收，会阻塞等待里面的done发出数据，会导致10个worker的执行不是并发的，顺序执行

		// 每个任务+1
		wg.Add(1)
	}
	wg.Wait()

	// for _, v := range workers {
	// 	<-v.done
	// }

}

func main() {
	chanDemo()
}

// 增加done chan 用来通知外面，已经处理完毕
func doWork(id int, c chan int, wg *sync.WaitGroup, done func()) {
	for {
		rlt, ok := <-c // 用两个值进行收

		// 如果没有收到数据，就进入下次循环
		if !ok {
			break
		}
		fmt.Printf("id%d接收到的是%d\n", id, rlt)
		// done <- true 这里如果直接发，外面没有接收的，会死锁

		done()
	}

}

// 创建worker，利用闭包，每个协程固定数据
func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
		done: func() {
			wg.Done()
		},
	}

	go doWork(id, w.in, wg, w.done)

	return w
}
