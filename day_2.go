// Solution to https://www.hackerrank.com/challenges/30-operators.
package main

import (
	"fmt"
	"math"
)

// Go doesn't know how to round numbers out of the box. Uh.
// Copied from https://github.com/golang/go/issues/4594#issuecomment-135336012.
func round(f float64) int {
	if math.Abs(f) < 0.5 {
		return 0
	}
	return int(f + math.Copysign(0.5, f))
}

func main() {
	var price, taxPercent, tipPercent float64
	fmt.Scanf("%v\n%v\n%v\n", &price, &tipPercent, &taxPercent)
	tax := price * taxPercent / 100
	tip := price * tipPercent / 100
	price += tax + tip
	fmt.Printf("The total meal cost is %v dollars.", round(price))
}
