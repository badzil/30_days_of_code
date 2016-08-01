// Solution to https://www.hackerrank.com/challenges/30-binary-trees.

package main

import (
	"errors"
	"fmt"
)

// Queue

type Queue struct {
	memory []*Node
}

func NewQueue() Queue {
	memory := make([]*Node, 0)
	return Queue{memory}
}

func (queue *Queue) Enqueue(node *Node) {
	queue.memory = append(queue.memory, node)
}

func (queue *Queue) Dequeue() (*Node, error) {
	if len(queue.memory) == 0 {
		// Queue is empty.
		return nil, errors.New("Queue is empty.")
	}
	node := queue.memory[0]
	queue.memory = queue.memory[1:]
	return node, nil
}

// Binary search tree

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(data int) Node {
	return Node{data: data}
}

func NewTree(datas []int) Node {
	root := NewNode(datas[0])
	for _, data := range datas[1:] {
		root.Insert(data)
	}
	return root
}

func (node *Node) Insert(data int) {
	if data <= node.data {
		if node.left == nil {
			newNode := NewNode(data)
			node.left = &newNode
		} else {
			node.left.Insert(data)
		}
	} else {
		if node.right == nil {
			newNode := NewNode(data)
			node.right = &newNode
		} else {
			node.right.Insert(data)
		}
	}
}

func (node *Node) LevelOrder(queue *Queue) {
	fmt.Printf("%v ", node.data)
	if node.left != nil {
		queue.Enqueue(node.left)
	}
	if node.right != nil {
		queue.Enqueue(node.right)
	}
	nextNode, err := queue.Dequeue()
	if err == nil {
		nextNode.LevelOrder(queue)
	}
}

func main() {
	datas := []int{3, 5, 4, 7, 2, 1}
	root := NewTree(datas)
	queue := NewQueue()
	root.LevelOrder(&queue)
	fmt.Println()
}
