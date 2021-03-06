// Solution to https://www.hackerrank.com/challenges/30-class-vs-instance

package main

import "fmt"

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	if initialAge < 0 {
		initialAge = 0
		fmt.Println("Age is not valid, setting age to 0.")
	}
	p.age = initialAge
	return p
}

func (p person) AmIOld() {
	if p.age < 13 {
		fmt.Println("You are young.")
	} else if 13 <= p.age && p.age < 18 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are old.")
	}

}

func (p person) YearPasses() person {
	p.age += 1
	return p
}

func main() {
	var T, age int

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&age)
		p := person{age: age}
		p = p.NewPerson(age)
		p.AmIOld()
		for j := 0; j < 3; j++ {
			p = p.YearPasses()
		}
		p.AmIOld()
		fmt.Println()
	}
}
