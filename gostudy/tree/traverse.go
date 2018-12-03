package tree

func (node *Node) Traverse() {
	if node == nil {
		return
	}

	//node.Left.Traverse()
	//node.Print()
	//node.Right.Traverse()

	node.TraverseFunc(func(n *Node) {
		// TraverseFunc, 此处传入的一个匿名函数，函数体进行node 节点的打印
		n.Print()
	})
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
