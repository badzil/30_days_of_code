// Solution to https://www.hackerrank.com/challenges/30-arrays.

package main

import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"

func reverseNumbers(numbers []int) []int {
	// WTF is wrong with Go??
	reversedNumbers := make([]int, len(numbers))
	for i := range numbers {
		reversedNumbers[i] = numbers[len(numbers)-i-1]
	}
	return reversedNumbers
}

func printlnNumbers(numbers []int) {
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
	array_size_str, _ := buffer.ReadString('\n')
	array_size, _ := strconv.Atoi(strings.TrimSpace(array_size_str))

	numbers := make([]int, array_size)
	for i := 0; i < array_size; i++ {
		number_str, _ := buffer.ReadString(' ')
		number, _ := strconv.Atoi(strings.TrimSpace(number_str))
		numbers[i] = number
	}

	printlnNumbers(reverseNumbers(numbers))
}
