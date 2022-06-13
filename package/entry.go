package main

import (
	"fmt"
	"learn/package/tree"
)

func main() {
	var root tree.TreeNode = tree.TreeNode{Value: 3}

	//root = treeNode{value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{1, nil, nil} // 不写键值，会根据结构体定义时的顺序进行赋值

	root.Right.Left = new(tree.TreeNode) // new函数开辟内存，返回地址， 指针也可以往下.

	nodes := []tree.TreeNode{
		{Value: 4},
		{},
	}

	root.Left.Left = tree.CreateNode(10)

	fmt.Println(root.Value, nodes)

	root.Print() // 和定义一个print函数，然后print函数第一个参数接收treeNode 是等价的
	tree.Print(root)

	root.SetValue(4)
	root.Print()

	proot := &root
	proot.SetVal(5) // setVal设置的是treeNode的指针，需要用指针调用
	root.Print()

	// nil指针也可以调用方法
	var pRoot *tree.TreeNode
	pRoot.SetVal(200)

	root.Tra()
}
