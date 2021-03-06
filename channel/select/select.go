package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var c1, c2 chan int = generator(), generator()
	work := createWorker(0)
	// c1,c2两个管道，谁先接收到值先打印谁
	//var hasVal bool = false

	var values []int
	n := 0
	//十秒后会向tm中发送一个数
	tm := time.After(10 * time.Second)

	// 定时任务, 每隔一段时间向chan中发一个值
	tick := time.Tick(time.Second)
	for {
		// active每次需要初始化

		var active chan int
		var activeValue int
		//if hasVal {
		//	active = work
		//}
		if len(values) > 0 {
			activeValue = values[0]
			active = work
		}

		/**
		 * 前两个case用来发送数据，第三个case用来接收数据
		 * 如果发送的过快，消耗的过慢，会打印不出来, 为了解决这个问题，需要建立一个存储，来存储没有打印的值
		 **/
		select {
		// 如果没有default, 只有select case，在没有传入值的时候，会死锁
		case n = <-c1:
			fmt.Println("c1", n)
			values = append(values, n)
			//hasVal = true
			//work <- n

		case n = <-c2:
			fmt.Println("c2", n)
			values = append(values, n)
			//hasVal = true
			//work <- n
			//		case active <- n: // 增加这行的意义：防止worker阻塞等待，但是n设置了初始值，会一直走着，利用变量来控制开关
			//			hasVal = false
		case active <- activeValue:
			//hasVal = true
			values = values[1:] // 去除第一个value
		case <-tm: // 模拟10s后结束
			fmt.Println("end")
			return
		case <-time.After(800 * time.Millisecond): // 如果两次生成数据相差了超过800毫秒，其他case能执行的时候，这个不会被执行
			fmt.Println("timeout")
		case <-tick: // 每秒查看队列长度
			fmt.Println("队列长度", len(values))
			//default:
			// 非阻塞式的获取channel，如果channel还没值，会走这，即使channel刚建立，还没有传入值，如果外部套for  这里会死循环
			//	fmt.Println("no")
			//	time.Sleep(time.Second)
		}
	}
}

func generator() chan int {
	out := make(chan int)

	fmt.Println(1)
	// 这个闭包会异步执行，先返回out
	go func() {
		fmt.Println(2)
		var i int = 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			out <- i
			i++
		}
	}()
	fmt.Println(3, &out)
	return out
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("worker %d 接收到%d\n", id, n)
	}
}
