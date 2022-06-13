/*
 * :Date: 2022-06-13 21:14:18
 * :LastEditTime: 2022-06-13 21:26:40
 * :Description:
 */
package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

/**
工厂函数，用来生成treeNode
**/
func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value} // c++中，在函数内部返回一个地址（返回局部变量地址），会报错。 go的内存分配会根据使用分配到堆上或者栈上。分配到栈上，在函数退出后会回收，堆上不会。c++上需要自己分配到堆上，自己回收，
}

/**
 * 第一个()里为接收者，为结构体，有this的意思, 接收者也为值传值
**/
func (node TreeNode) Print() {
	fmt.Println(node.Value)
}

func Print(node TreeNode) {
	fmt.Println(node.Value)
}

/**
 * 函数里设置node的值，不会改变外部的
**/
func (node TreeNode) SetValue(val int) {
	node.Value = val
}

/**
 * 可以改变值，使用的是treeNode的指针
**/
func (node *TreeNode) SetVal(val int) {
	if node == nil {
		fmt.Println("指针为空")
		return
	}
	node.Value = val
}

/**
  *
**/
func (node *TreeNode) Tra() {
	if node == nil {
		return
	}
	// 深度优先遍历
	node.Left.Tra()
	node.Print()
	node.Right.Tra()
}

/**
 * 指针接收者和值接受者使用
 * 如果要改变内容，使用指针接收者
 * 结构过大考虑使用指针接收者
 * 一致性：如果有指针接收者，最好都是指针接收者
 * 值接收者是go 特有
 *
**/
