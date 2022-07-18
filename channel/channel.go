package main

// https://blog.csdn.net/tool007/article/details/124329558  参考
import (
	"fmt"
	"time"
)

// channel使用

// 无缓冲channel
func chanDemo() {
	var c chan int     // 定义channel里面存int类型
	c = make(chan int) // 无缓存管道，没有被消费前，不能写入数据，会死锁all goroutines are asleep - deadlock

	//c <- 3
	go worker(10, c)
	go func(c chan int) {
		for {
			n := <-c // 从c里取一个
			fmt.Println(n)
		}
	}(c)

	c <- 1 // 在管道里存入1

	c <- 2 // 在管道里存入2
	//c <- 2 // 在管道里存入2

	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- i
	}

	var cc [10]chan<- int
	for i := 10; i < 20; i++ {
		cc[i-10] = createWorker(i)
		// 给每个建立好的channel发送其所需数据
		cc[i-10] <- i
	}

}

// 有缓冲channel
func bufferedChannel() {
	c := make(chan int, 3) // 每次向channel发送数据或者从channel读取数据都会发生协程的切换，频繁切换对性能也有影响, 可以往channel里存三个数据，不发生协程切换
	c <- 1
	c <- 2
	c <- 3
	go worker(1, c)

}

func main() {
	//chanDemo()
	bufferedChannel()
	channelClose()
	time.Sleep(time.Second) // 防止还没运行，就退出了
}

func worker(id int, c chan int) {
	for {
		rlt, ok := <-c // 用两个值进行收

		// 如果没有收到数据，就进入下次循环
		if !ok {
			break
		}
		fmt.Printf("id%d接收到的是%d\n", id, rlt)
	}

	// 方法二
	// for rlt :=range c {
	// 	fmt.Printf("id%d接收到的是%d\n", id, rlt)
	// }
}

// 创建worker，等待数据接收
func createWorker(id int) chan<- int {
	c := make(chan int)
	// 闭包里写每个worker的实现，每个建立一个channel，用来接收自己的数据
	go func() {
		for {
			fmt.Printf("create id%d接收到的是%d\n", id, <-c)
		}
	}()
	return c
}

// channel 数据发送方，可以通过close来说明，数据发送完成，只有发送方才能close

func channelClose() {
	c := make(chan int)
	go worker(2, c)
	c <- 1
	c <- 2
	close(c) // 关闭channel, 关闭后，接受方辉一直收0，直到程序退出, 一般不需要使用
}
