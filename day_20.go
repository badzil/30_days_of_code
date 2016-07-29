// Solution to https://www.hackerrank.com/challenges/30-sorting.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type SortableArray struct {
	nums     []int
	numSwaps int
}

func (array *SortableArray) Swap(idx1, idx2 int) {
	array.nums[idx1], array.nums[idx2] = array.nums[idx2], array.nums[idx1]
	array.numSwaps += 1
}

func (array *SortableArray) BubbleSort() {
	array.numSwaps = 0
	for idx1 := 0; idx1 < len(array.nums); idx1++ {
		for idx2 := idx1; idx2 < len(array.nums); idx2++ {
			if array.nums[idx1] > array.nums[idx2] {
				array.Swap(idx1, idx2)
			}
		}
	}
}

func GetNumbers() []int {
	buffer := bufio.NewReader(os.Stdin)
	// Get rid of the number of numbers.
	buffer.ReadString('\n')
	line, _ := buffer.ReadString('\n')
	tokens := strings.Split(line, " ")
	nums := make([]int, len(tokens))
	for index, token := range tokens {
		nums[index] = SafeAtoi(token)
	}
	return nums
}

func SafeAtoi(str string) int {
	integer, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return integer
}

func SortNumbers(nums []int) {
	array := SortableArray{nums: nums}
	array.BubbleSort()
	fmt.Printf("Array is sorted in %v swaps.\n", array.numSwaps)
	fmt.Println("First Element:", array.nums[0])
	fmt.Println("Last Element:", array.nums[len(array.nums)-1])
}

func main() {
	nums := GetNumbers()
	SortNumbers(nums)
}
