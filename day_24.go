// Solution to https://www.hackerrank.com/challenges/30-linked-list-deletion.

package main

import "fmt"

type Node struct {
	data      int
	childNode *Node
}

func (parentNode *Node) Display() {
	for parentNode != nil {
		fmt.Printf("%v ", parentNode.data)
		parentNode = parentNode.childNode
	}
	fmt.Println()
}

func (parentNode *Node) DeleteDuplicates() {
	childNode := parentNode.childNode
	if childNode != nil {
		if parentNode.data == childNode.data {
			parentNode.childNode = childNode.childNode
		}
		childNode.DeleteDuplicates()
	}
}

func Insert(head *Node, data int) *Node {
	if head == nil {
		head = &Node{data: data}
	} else {
		parentNode := head
		for parentNode.childNode != nil {
			parentNode = parentNode.childNode
		}
		parentNode.childNode = &Node{data: data}
	}
	return head
}

func main() {
	var head *Node
	for _, number := range []int{1, 2, 2, 3, 3, 4} {
		head = Insert(head, number)
	}
	head.DeleteDuplicates()
	head.Display()
}
