// Solution to https://www.hackerrank.com/challenges/30-queues-stacks.

package main

import (
	"fmt"
)

// Stack

type Stack struct {
	memory []rune
}

func NewStack() Stack {
	memory := make([]rune, 0)
	return Stack{memory}
}

func (stack *Stack) PushCharacter(char rune) {
	stack.memory = append(stack.memory, char)
}

func (stack *Stack) PopCharacter() rune {
	char := stack.memory[len(stack.memory)-1]
	stack.memory = stack.memory[:len(stack.memory)-1]
	return char
}

// Queue

type Queue struct {
	memory []rune
}

func NewQueue() Queue {
	memory := make([]rune, 0)
	return Queue{memory}
}

func (queue *Queue) EnqueueCharacter(char rune) {
	queue.memory = append(queue.memory, char)
}

func (queue *Queue) DequeueCharacter() rune {
	char := queue.memory[0]
	queue.memory = queue.memory[1:]
	return char
}

func IsPalindrome(input string) bool {
	stack, queue := NewStack(), NewQueue()
	for _, char := range input {
		stack.PushCharacter(char)
		queue.EnqueueCharacter(char)
	}
	for i := 0; i < len(input); i++ {
		startChar := queue.DequeueCharacter()
		endChar := stack.PopCharacter()
		if startChar != endChar {
			return false
		}
	}
	return true
}

func TestStack() {
	s := NewStack()
	s.PushCharacter('a')
	s.PushCharacter('b')
	fmt.Println(string(s.PopCharacter()))
	fmt.Println(string(s.PopCharacter()))
}

func TestQueue() {
	q := NewQueue()
	q.EnqueueCharacter('a')
	q.EnqueueCharacter('b')
	fmt.Println(string(q.DequeueCharacter()))
	fmt.Println(string(q.DequeueCharacter()))
}

func main() {
	// TestStack()
	// TestQueue()

	for _, input := range []string{"racecar", "susan"} {
		var modifier string
		if !IsPalindrome(input) {
			modifier = "not "
		}
		fmt.Printf("The word, %v, is %va palindrome.\n", input, modifier)
	}
}
