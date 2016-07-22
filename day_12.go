// Solution to https://www.hackerrank.com/challenges/30-inheritance.

package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	idNumber  int
}

type Student struct {
	Person
	scores []int
}

func (person Person) PrintPerson() {
	fmt.Printf(
		"Name: %v, %v\nID: %v\n",
		person.lastName,
		person.firstName,
		person.idNumber,
	)
}

func (student Student) Calculate() (letter string) {
	average := GetAverage(student.scores)
	return GetGradeCharacter(average)
}

func GetAverage(numbers []int) float64 {
	sum := 0
	for _, score := range numbers {
		sum += score
	}
	return float64(sum) / float64(len(numbers))
}

func GetGradeCharacter(average float64) string {
	/* Convert a score average to a grade character. */
	letter := "N/A"
	switch {
	case average < 40:
		letter = "T"
	case 40 <= average && average < 55:
		letter = "D"
	case 55 <= average && average < 70:
		letter = "P"
	case 70 <= average && average < 80:
		letter = "A"
	case 80 <= average && average < 90:
		letter = "E"
	case 90 <= average && average <= 100:
		letter = "O"
	}
	return letter
}

func main() {
	s := Student{Person{"Heraldo", "Memelli", 8135627}, []int{100, 80}}
	s.PrintPerson()
	fmt.Printf("Grade: %v\n", s.Calculate())
}
