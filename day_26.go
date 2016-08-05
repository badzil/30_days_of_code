// Solution to https://www.hackerrank.com/challenges/30-nested-logic.

package main

import (
	"fmt"
	"os"
)

func GetFine(eyear, emonth, eday, ayear, amonth, aday int) int {
	if ayear == eyear {
		if amonth == emonth && aday > eday {
			return 15 * (aday - eday)
		} else if amonth > emonth {
			return 500 * (amonth - emonth)
		}
	} else if ayear > eyear {
		return 10000
	}
	return 0
}

func RunTests() {
	AssertEqual("1 day early", GetFine(2016, 8, 2, 2016, 8, 1), 0)
	AssertEqual("1 day early over month break", GetFine(2016, 2, 1, 2016, 1, 31), 0)
	AssertEqual("1 day early over year break", GetFine(2017, 1, 1, 2016, 12, 31), 0)
	AssertEqual("Same day", GetFine(2016, 8, 2, 2016, 8, 2), 0)
	AssertEqual("1 day late", GetFine(2016, 8, 2, 2016, 8, 3), 15)
	AssertEqual("1 day late over month break", GetFine(2016, 7, 31, 2016, 8, 1), 500)
	AssertEqual("1 day late over year break", GetFine(2016, 12, 31, 2017, 1, 1), 10000)
	AssertEqual("2 days late", GetFine(2016, 8, 2, 2016, 8, 4), 30)
	AssertEqual("2 days late over month break", GetFine(2016, 7, 31, 2016, 8, 2), 500)
	AssertEqual("1 month late", GetFine(2016, 8, 2, 2016, 9, 2), 500)
	AssertEqual("1 month late over year break", GetFine(2016, 12, 2, 2017, 1, 2), 10000)
}

func AssertEqual(msg string, left, right interface{}) {
	if left != right {
		fmt.Printf("[%v] %v != %v", msg, left, right)
		os.Exit(1)
	}
}

func ReadDate() (year, month, day int) {
	fmt.Scanf("%v %v %v", &day, &month, &year)
	return year, month, day
}

func main() {
	RunTests()

	ayear, amonth, aday := ReadDate()
	eyear, emonth, eday := ReadDate()
	fmt.Println(GetFine(eyear, emonth, eday, ayear, amonth, aday))
}
