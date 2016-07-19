// Solution to https://www.hackerrank.com/challenges/30-binary-numbers.

package main

import "fmt"

func GetInput() int {
	var n int
	fmt.Scanf("%v", &n)
	return n
}

func AssertEqual(a int, b int) {
	if a == b {
		fmt.Print(".")
	} else {
		fmt.Printf("Error: %v != %v", a, b)
	}
}

func DivWithRemainder(dividend int, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend - quotient*divisor
	return quotient, remainder
}

func GetConsecutives(n int) int {
	bit, consecutives, currentConsecutives := 0, 0, 0
	for n != 0 {
		n, bit = DivWithRemainder(n, 2)
		if bit == 0 {
			if currentConsecutives > consecutives {
				consecutives = currentConsecutives
			}
			currentConsecutives = 0
		} else {
			currentConsecutives++
		}
	}
	if currentConsecutives > consecutives {
		consecutives = currentConsecutives
	}
	return consecutives
}

func RunTests() {
	var quotient int
	var remainder int
	quotient, remainder = DivWithRemainder(2, 2)
	AssertEqual(quotient, 1)
	AssertEqual(remainder, 0)
	quotient, remainder = DivWithRemainder(5, 2)
	AssertEqual(quotient, 2)
	AssertEqual(remainder, 1)

	AssertEqual(GetConsecutives(0), 0)  // 000
	AssertEqual(GetConsecutives(1), 1)  // 001
	AssertEqual(GetConsecutives(2), 1)  // 010
	AssertEqual(GetConsecutives(3), 2)  // 011
	AssertEqual(GetConsecutives(4), 1)  // 100
	AssertEqual(GetConsecutives(5), 1)  // 101
	AssertEqual(GetConsecutives(6), 2)  // 110
	AssertEqual(GetConsecutives(7), 3)  // 111
	AssertEqual(GetConsecutives(27), 2) // 11011
}

func main() {
	//RunTests()
	n := GetInput()
	fmt.Println(GetConsecutives(n))
}
