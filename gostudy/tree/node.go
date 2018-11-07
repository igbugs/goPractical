package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) Setvalue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignored.")
		return // 如果此处不进行return, 下面nil.Value 是取不到值的，会进行报错
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
	// 此处的 &Node 为函数的局部变量的指针，在其他的语言中，返回局部变量给其他人使用，是不可以的，go语言可行
	// 工厂函数一般返回 重新构造的结构体的地址
}
