package main

import (
	"tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}           // 定义了root 的值为3，没有子节点
	root.Left = &tree.Node{}             // root 的左节点 指向的子节点的地址
	root.Right = &tree.Node{5, nil, nil} // 赋值root的右节点，为5 没有字节点（左右节点指针为nil）
	root.Right.Left = new(tree.Node)     // root.Right 为一个指针，指针可以直接取Left属性，go语言会进行变量和指针的相互转换
	root.Left.Right = tree.CreateNode(2)

	//nodes := []Node {			// 定义了TreeNode 的切片，含有三个的节点，之间并没有联系
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//
	//fmt.Println(nodes)
	//fmt.Println(root)

	//root.Print()
	//fmt.Println()
	//root.Right.Print()

	//fmt.Println()
	root.Right.Left.Setvalue(4)
	//root.Right.Left.Print()

	//var pRoot *Node       // 刚初始化，pRoot 这个结构体地址为nil
	//pRoot.Setvalue(200)
	//pRoot = &root
	//pRoot.Setvalue(300)
	//pRoot.Print()

	root.Traverse()

	//fmt.Println()
	//myRoot := myTreeNode{&root}
	//myRoot.postOrder()
	fmt.Println()

	nodeCount := 0
	root.TraverseFunc(func(n *tree.Node) {
		nodeCount++
	})

	fmt.Println("Node count: ", nodeCount)
}
