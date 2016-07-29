// Solution to https://www.hackerrank.com/challenges/30-interfaces.

package main

import (
	"fmt"
	"os"
	"reflect"
)

type AdvancedArithmetic interface {
	DivisorSum(n int) int
}

type Calculator struct {
}

func (calc Calculator) DivisorSum(n int) int {
	divisorSum := 1
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			divisorSum += i
		}
	}
	return divisorSum
}

func DivisorSumWithInterface(aa AdvancedArithmetic, n int) int {
	/* This is solely to try to replicate the Hackerrank problem, and
	   verify that I can figure out how to work with interfaces in Go. */
	fmt.Println("I am a:", reflect.TypeOf(aa))
	return aa.DivisorSum(n)
}

func main() {
	calc := Calculator{}
	fmt.Println(DivisorSumWithInterface(calc, 6))
}
