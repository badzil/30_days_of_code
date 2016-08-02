// Solution to https://www.hackerrank.com/challenges/30-linked-list-deletion.

package main

import "./linkedlist"

func DeleteDuplicates(parentNode *linkedlist.Node) {
	childNode := parentNode.ChildNode
	if childNode != nil {
		if parentNode.Data == childNode.Data {
			parentNode.ChildNode = childNode.ChildNode
		}
		DeleteDuplicates(childNode)
	}
}

func main() {
	var head *linkedlist.Node
	for _, number := range []int{1, 2, 2, 3, 3, 4} {
		head = linkedlist.Insert(head, number)
	}
	DeleteDuplicates(head)
	head.Display()
}
