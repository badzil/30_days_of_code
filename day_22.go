// Solution to https://www.hackerrank.com/challenges/30-binary-search-trees.

package main

import "fmt"

type Node struct {
    data int
    left *Node
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

func (node *Node) GetHeight() int {
    leftHeight, rightHeight := 0, 0
    if node.left != nil {
        leftHeight = node.left.GetHeight()
    }
    if node.right != nil {
        rightHeight = node.right.GetHeight()
    }
    return 1 + Max(leftHeight, rightHeight)
}

func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func main() {
    datas := []int{3, 5, 2, 1, 4, 6, 7}
    root := NewTree(datas)
    // GetHeight counts the number of nodes, but the expected answer is
    // number of edges. Therefore we have to remove 1.
    fmt.Println(root.GetHeight() - 1)
}
