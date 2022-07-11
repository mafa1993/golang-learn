package queue

import "fmt"

func ExampleQueue_Pop() {
	q := Queue{}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	// Output:
	//2
	//3
}
