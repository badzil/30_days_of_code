// Solution to https://www.hackerrank.com/challenges/30-dictionaries-and-maps.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func safeAtoi(str string) int {
	integer, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return integer
}

func getProblemSize(buffer *bufio.Reader) int {
	problem_size_str, _ := buffer.ReadString('\n')
	return safeAtoi(problem_size_str)
}

func getPhoneNumbers(buffer *bufio.Reader) map[string]int {
	problem_size := getProblemSize(buffer)
	phone_numbers := make(map[string]int)
	for i := 0; i < problem_size; i++ {
		entry, _ := buffer.ReadString('\n')
		tokens := strings.Split(entry, " ")
		name, phone_number_str := tokens[0], tokens[1]
		phone_numbers[name] = safeAtoi(phone_number_str)
	}
	return phone_numbers
}

func printPhoneNumber(phone_numbers map[string]int, name string) {
	phone_number, ok := phone_numbers[name]
	if ok {
		fmt.Printf("%v=%v\n", name, phone_number)
	} else {
		fmt.Println("Not found")
	}
}

func main() {
	buffer := bufio.NewReader(os.Stdin)
	phone_numbers := getPhoneNumbers(buffer)

	for {
		name, err := buffer.ReadString('\n')
		printPhoneNumber(phone_numbers, strings.TrimSpace(name))
		if err != nil {
			// EOF
			break
		}
	}
}
