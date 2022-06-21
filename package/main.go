/*
 * :Date: 2022-06-21 20:43:25
 * :LastEditTime: 2022-06-21 20:55:03
 * :Description:
 */
package main

import (
	"fmt"
	"learn/package/queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
