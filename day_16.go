// Solution to https://www.hackerrank.com/challenges/30-exceptions-string-to-integer.

package main

import (
	"fmt"
	"strconv"
)

func CastAndPrintInt(s string) {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Bad String")
	} else {
		fmt.Println(i)
	}
}

func main() {
	CastAndPrintInt("3")
	CastAndPrintInt("za")
}
