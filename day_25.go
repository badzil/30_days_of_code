// Solution to https://www.hackerrank.com/challenges/30-running-time-and-complexity.

package main

import (
	"fmt"
	"math"
	"os"
)

func IsPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i <= IntSqrt(n); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IntSqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}

func AssertEqual(msg string, left, right interface{}) {
	if left != right {
		fmt.Printf("[%v] %v != %v\n", msg, left, right)
		os.Exit(1)
	}
}

func main() {
	AssertEqual("1", IsPrime(1), false)
	AssertEqual("2", IsPrime(2), true)
	AssertEqual("3", IsPrime(3), true)
	AssertEqual("4", IsPrime(4), false)
	AssertEqual("5", IsPrime(5), true)
	AssertEqual("6", IsPrime(6), false)
	AssertEqual("7", IsPrime(7), true)
	AssertEqual("8", IsPrime(8), false)
	AssertEqual("9", IsPrime(9), false)
	AssertEqual("10", IsPrime(10), false)
	AssertEqual("11", IsPrime(11), true)
	AssertEqual("12", IsPrime(12), false)

	var n, input int
	fmt.Scanf("%v", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v", &input)
		if IsPrime(input) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}
	}
}
