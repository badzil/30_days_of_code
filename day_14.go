package main

import (
	"fmt"
	"math"
)

type Difference struct {
	elements          []int
	maximumDifference int
}

func NewDifference(elements []int) Difference {
	defaultMaximumDifference := 0
	return Difference{elements, defaultMaximumDifference}
}

func (diff *Difference) ComputeDifference() {
	for idx, a := range diff.elements {
		for _, b := range diff.elements[idx+1:] {
			currentDifference := GetDifference(a, b)
			if currentDifference > diff.maximumDifference {
				diff.maximumDifference = currentDifference
			}
		}
	}
}

func GetDifference(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func main() {
	elements := []int{1, 2, 5}
	d := NewDifference(elements)
	d.ComputeDifference()
	fmt.Println(d.maximumDifference)
}
