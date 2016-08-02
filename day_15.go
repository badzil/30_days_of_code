// Solution to https://www.hackerrank.com/challenges/30-linked-list.

package main

import "./linkedlist"

func main() {
	var head *linkedlist.Node
	for _, number := range []int{2, 3, 4, 1} {
		head = linkedlist.Insert(head, number)
	}
	head.Display()
}
