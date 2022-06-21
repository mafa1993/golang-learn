/*
 * :Date: 2022-06-21 20:38:25
 * :LastEditTime: 2022-06-21 20:48:10
 * :Description:
 */
package queue

type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0] // 取出头部元素
	*q = (*q)[1:]     //改变q的值
	return head
}

// 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
