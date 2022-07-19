package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

func (n *node) tra(f func(*node)) {
	if n == nil {
		return
	}
	n.left.tra(f)
	fmt.Println("前")
	fmt.Println("内", n.val)
	f(n)
	fmt.Println("后")
	n.right.tra(f)
}

func main() {
	n := &node{val: 0}
	//n.left = &node{val: 1}
	n.left = &node{
		val: 1,
		left: &node{
			val: 2,
		},
	}

	n.right = &node{
		val:  3,
		left: &node{},
	}
	n.right.left.val = 4
	c := n.traWithChan()
	for v := range c {
		fmt.Println("外", v.val)
	}
}

func (n *node) traWithChan() chan *node {
	c := make(chan *node)
	// 执行后，放入node后，会阻塞等待有人从管道c中接收走
	go func() {
		n.tra(func(node *node) {
			fmt.Println("c", &c)
			c <- node // 第一轮不会阻塞，第二轮在这阻塞等待，等到channel中的值被接收走，在向channel中放第二个值之前，被接收走即可
			fmt.Println("cc", c)
		})
		close(c)
	}()
	return c
}
