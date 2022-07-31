/*
 * :Date: 2022-06-13 21:11:28
 * :LastEditTime: 2022-06-21 20:36:56
 * :Description:
 */
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
	
	//使用组合对已有结构体扩展
	myRoot := myTree{&root}
	myRoot.fore()
}

// 使用别名，封装tree
type myTree struct {
	node *tree.TreeNode
}

// TreeNode是先遍历的左，又遍历的右，这里改为先遍历右，再遍历左
func (myNode *myTree) fore() {
	if myNode == nil || myNode.node == nil {
		return
	}
	right := myTree{myNode.node.Right} // 类型转换为myTree，不软不能调用fore
	left := myTree{myNode.node.Left}
	right.fore()  // 包一层进行循环
	left.fore()
	myNode.node.Print()
}
