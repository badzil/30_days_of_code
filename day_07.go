// Solution to https://www.hackerrank.com/challenges/30-arrays.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReverseNumbers(numbers []int) []int {
	// WTF is wrong with Go??
	reversedNumbers := make([]int, len(numbers))
	for i := range numbers {
		reversedNumbers[i] = numbers[len(numbers)-i-1]
	}
	return reversedNumbers
}

func PrintlnNumbers(numbers []int) {
	for index, number := range numbers {
		if index != 0 {
			fmt.Print(" ")
		}
		fmt.Print(number)
	}
	fmt.Print("\n")
}

func main() {
	buffer := bufio.NewReader(os.Stdin)
	arraySizeStr, _ := buffer.ReadString('\n')
	arraySize, _ := strconv.Atoi(strings.TrimSpace(arraySizeStr))

	numbers := make([]int, arraySize)
	for i := 0; i < arraySize; i++ {
		numberStr, _ := buffer.ReadString(' ')
		number, _ := strconv.Atoi(strings.TrimSpace(numberStr))
		numbers[i] = number
	}

	PrintlnNumbers(ReverseNumbers(numbers))
}
