package linkedlist

import "fmt"

type Node struct {
	Data      int
	ChildNode *Node
}

func (parentNode *Node) Display() {
	for parentNode != nil {
		fmt.Printf("%v ", parentNode.Data)
		parentNode = parentNode.ChildNode
	}
	fmt.Println()
}

func Insert(head *Node, data int) *Node {
	if head == nil {
		head = &Node{Data: data}
	} else {
		parentNode := head
		for parentNode.ChildNode != nil {
			parentNode = parentNode.ChildNode
		}
		parentNode.ChildNode = &Node{Data: data}
	}
	return head
}
