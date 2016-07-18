// Solution to https://www.hackerrank.com/challenges/30-conditional-statements.

package main

import "fmt"

func IsWeird(number int) bool {
	return (6 <= number && number <= 20) || number%2 == 1
}

func main() {
	var number int
	fmt.Scanf("%v", &number)
	if IsWeird(number) {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}

}
