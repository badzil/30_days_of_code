// Solution to https://www.hackerrank.com/challenges/30-more-exceptions.

package main

import (
	"errors"
	"fmt"
)

func ComputePower(n, power int) (int, error) {
	if n < 0 || power < 0 {
		err := errors.New("n and p should be non-negative")
		return 0, err
	}
	result := 1
	if power != 0 {
		for i := 0; i < power; i++ {
			result *= n
		}
	}
	return result, nil
}

func PrintPower(n, power int) {
	result, err := ComputePower(n, power)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func main() {
	// Ideal cases
	PrintPower(0, 0)
	PrintPower(1, 0)
	PrintPower(0, 1)
	PrintPower(3, 5)
	PrintPower(2, 4)

	// Problem cases
	PrintPower(-1, -2)
	PrintPower(-1, 3)
}
