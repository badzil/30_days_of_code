// Solution to https://www.hackerrank.com/challenges/30-regex-patterns.

package main

import (
	"fmt"
	"sort"
	"strings"
)

type Person struct {
	Name  string
	Email string
}

type ByName []Person

func (e ByName) Len() int           { return len(e) }
func (e ByName) Less(i, j int) bool { return e[i].Name < e[j].Name }
func (e ByName) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

func (p Person) GetEmailDomain() string {
	return strings.Split(p.Email, "@")[1]
}

func GetPeople() []Person {
	var n int
	fmt.Scanf("%v", &n)
	people := make([]Person, n)
	for i := 0; i < n; i++ {
		var name, email string
		fmt.Scanf("%v %v", &name, &email)
		people[i] = Person{Name: name, Email: email}
	}
	return people
}

func main() {
	people := GetPeople()
	sort.Sort(ByName(people))
	for _, person := range people {
		if person.GetEmailDomain() == "gmail.com" {
			fmt.Println(person.Name)
		}
	}
}
