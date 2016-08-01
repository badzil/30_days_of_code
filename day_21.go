// Solution to https://www.hackerrank.com/challenges/30-generics.

package main

import "fmt"
import "reflect"

func PrintArray(input interface{}) {
	slice := reflect.ValueOf(input)
	for idx := 0; idx < slice.Len(); idx++ {
		fmt.Println(slice.Index(idx))
	}
}

func main() {
	intArray := []int{1, 2, 3}
	stringArray := []string{"Hello", "World"}
	PrintArray(intArray)
	PrintArray(stringArray)
}
