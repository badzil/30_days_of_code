// Solution to https://www.hackerrank.com/challenges/30-testing.

package main

import (
	"fmt"
	"os"
)

type TestCase struct {
	ArrivalTimes          []int
	CancellationThreshold int
}

func NewTestCase(arrivalTimes []int, cancellationThreshold int) TestCase {
	t := TestCase{
		ArrivalTimes:          arrivalTimes,
		CancellationThreshold: cancellationThreshold,
	}
	isValid, msg := t.IsValid()
	if !isValid {
		fmt.Printf("Test case %v is not valid: %v\n", t, msg)
		os.Exit(1)
	}
	return t
}

func (t TestCase) IsValid() (isValid bool, reason string) {
	numStudents := len(t.ArrivalTimes)
	if !(1 <= numStudents && numStudents <= 200 &&
		1 <= t.CancellationThreshold && t.CancellationThreshold <= numStudents) {
		return false, "Wrong number of students or cancellation threshold."
	}

	hasZero, hasPositive, hasNegative := false, false, false
	for _, arrivalTime := range t.ArrivalTimes {
		if arrivalTime < -1000 || 1000 < arrivalTime {
			return false, "Arrival time out of bounds."
		}
		if arrivalTime < 0 {
			hasNegative = true
		} else if arrivalTime == 0 {
			hasZero = true
		} else if arrivalTime > 0 {
			hasPositive = true
		}
	}
	if !hasNegative || !hasZero || !hasPositive {
		return false, "Missing one of negative, zero, positive."
	}
	return true, ""
}

func (t TestCase) Print() {
	fmt.Printf("%v %v\n", len(t.ArrivalTimes), t.CancellationThreshold)
	for _, i := range t.ArrivalTimes {
		fmt.Printf("%v ", i)
	}
	fmt.Println()
}

func IsCancelled(t TestCase) bool {
	numOnTime := 0
	for _, arrivalTime := range t.ArrivalTimes {
		if arrivalTime <= 0 {
			numOnTime += 1
			if numOnTime >= t.CancellationThreshold {
				return false
			}
		}
	}
	return true
}

func PrintTestCase(t TestCase, expectedResult bool) {
	AssertEqual(IsCancelled(t), expectedResult)
	t.Print()
}

func AssertEqual(left, right interface{}) {
	if left != right {
		fmt.Printf("%v != %v", left, right)
		os.Exit(1)
	}
}

func main() {
	fmt.Println(5)
	PrintTestCase(NewTestCase([]int{-1, 0, 1}, 3), true)
	PrintTestCase(NewTestCase([]int{-1, 0, 1}, 2), false)
	PrintTestCase(NewTestCase([]int{-2, -1, 0, 1}, 4), true)
	PrintTestCase(NewTestCase([]int{-2, -1, 0, 1}, 3), false)
	PrintTestCase(NewTestCase([]int{-2, -1, 0, 1, 2}, 4), true)
}
