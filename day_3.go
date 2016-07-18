// Solution to day 3.
package main

import (
	"fmt"
	"math"
)

func IsWeird(number float64) bool {
	return (6 <= number && number <= 20) || math.Mod(number, 2) == 1
}

func main() {
	var number float64
	fmt.Scanf("%v", &number)
	if IsWeird(number) {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}

}
