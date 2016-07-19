// Solution to https://www.hackerrank.com/challenges/30-recursion.

package main

import "fmt"

func Factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * Factorial(n-1)
	}
}

func main() {
	var n int
	fmt.Scanf("%v", &n)
	fmt.Println(Factorial(n))
}
