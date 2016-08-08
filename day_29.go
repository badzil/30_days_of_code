// Solution to https://www.hackerrank.com/challenges/30-bitwise-and.

package main

import (
	"fmt"
	"os"
)

func GetMaxAnd(size, threshold int) (maxAnd int) {
	for a := 1; a <= size; a++ {
		for b := a + 1; b <= size; b++ {
			res := a & b
			if res < threshold && res > maxAnd {
				maxAnd = res
			}
		}
	}
	return
}

func AssertEqual(left, right int) {
	if left != right {
		fmt.Printf("%v != %v\n", left, right)
		os.Exit(1)
	}
}

func RunTests() {
	AssertEqual(GetMaxAnd(5, 2), 1)
	AssertEqual(GetMaxAnd(8, 5), 4)
	AssertEqual(GetMaxAnd(2, 2), 0)
}

func main() {
	RunTests()

	var n int
	fmt.Scanf("%v", &n)
	for i := 0; i < n; i++ {
		var size, threshold int
		fmt.Scanf("%v %v", &size, &threshold)
		fmt.Println(GetMaxAnd(size, threshold))
	}

}
