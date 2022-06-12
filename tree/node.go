package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func main() {
	var root treeNode = treeNode{value: 3}

	//root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{1, nil, nil} // 不写键值，会根据结构体定义时的顺序进行赋值

	root.right.left = new(treeNode) // new函数开辟内存，返回地址， 指针也可以往下.

	nodes := []treeNode{
		{value: 4},
		{},
	}

	root.left.left = createNode(10)

	fmt.Println(root.value, nodes)

	root.print() // 和定义一个print函数，然后print函数第一个参数接收treeNode 是等价的
	print(root)

	root.setValue(4)
	root.print()

	proot := &root
	proot.setVal(5) // setVal设置的是treeNode的指针，需要用指针调用
	root.print()

	// nil指针也可以调用方法
	var pRoot *treeNode
	pRoot.setVal(200)

	root.tra()
}

/**
工厂函数，用来生成treeNode
**/
func createNode(value int) *treeNode {
	return &treeNode{value: value} // c++中，在函数内部返回一个地址（返回局部变量地址），会报错。 go的内存分配会根据使用分配到堆上或者栈上。分配到栈上，在函数退出后会回收，堆上不会。c++上需要自己分配到堆上，自己回收，
}

/**
 * 第一个()里为接收者，为结构体，有this的意思, 接收者也为值传值
**/
func (node treeNode) print() {
	fmt.Println(node.value)
}

func print(node treeNode) {
	fmt.Println(node.value)
}

/**
 * 函数里设置node的值，不会改变外部的
**/
func (node treeNode) setValue(val int) {
	node.value = val
}

/**
 * 可以改变值，使用的是treeNode的指针
**/
func (node *treeNode) setVal(val int) {
	if node == nil {
		fmt.Println("指针为空")
		return
	}
	node.value = val
}

/**
  *
**/
func (node *treeNode) tra() {
	if node == nil {
		return
	}
	// 深度优先遍历
	node.left.tra()
	node.print()
	node.right.tra()
}

/**
 * 指针接收者和值接受者使用
 * 如果要改变内容，使用指针接收者
 * 结构过大考虑使用指针接收者
 * 一致性：如果有指针接收者，最好都是指针接收者
 * 值接收者是go 特有
 * 
**/
