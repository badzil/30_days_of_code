// Solution to https://www.hackerrank.com/challenges/30-review-loop.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SafeAtoi(str string) int {
	integer, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return integer
}

func GetProblemSize(buffer *bufio.Reader) int {
	problemSizeStr, _ := buffer.ReadString('\n')
	return SafeAtoi(problemSizeStr)
}

func PrintOddAndEvenCharacters(word string) {
	oddChars, evenChars := "", ""
	for index, character := range word {
		if (index+1)%2 == 1 {
			oddChars += string(character)
		} else {
			evenChars += string(character)
		}
	}
	fmt.Printf("%v %v\n", oddChars, evenChars)
}

func main() {
	buffer := bufio.NewReader(os.Stdin)
	problemSize := GetProblemSize(buffer)
	for i := 0; i < problemSize; i++ {
		newline, _ := buffer.ReadString('\n')
		word := strings.TrimSpace(newline)
		PrintOddAndEvenCharacters(word)
	}
}
