// Solution to https://www.hackerrank.com/challenges/30-data-types.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInput() (int, float64, string) {
	reader := bufio.NewReader(os.Stdin)
	firstInput, _ := reader.ReadString('\n')
	secondInput, _ := reader.ReadString('\n')
	thirdInput, _ := reader.ReadString('\n')

	// Need to strip spaces otherwise convert methods fail.
	integer, _ := strconv.Atoi(strings.TrimSpace(firstInput))
	double, _ := strconv.ParseFloat(strings.TrimSpace(secondInput), 64)

	return integer, double, thirdInput
}

func main() {
	var firstInt, secondInt int
	var firstDouble, secondDouble float64
	var firstString, secondString string

	firstInt = 4
	firstDouble = 4.0
	firstString = "HackerRank "

	secondInt, secondDouble, secondString = GetInput()

	fmt.Println(firstInt + secondInt)
	fmt.Printf("%.1f\n", firstDouble+secondDouble)
	fmt.Println(firstString + secondString)
}
