// Solution to https://www.hackerrank.com/challenges/30-loops.

package main

import "fmt"

func PrintMultiplication(N, i int) {
	fmt.Printf("%v x %v = %v\n", N, i, N*i)
}

func main() {
	var N int
	fmt.Scanf("%v\n", &N)
	for i := 1; i <= 10; i++ {
		PrintMultiplication(N, i)
	}
}
