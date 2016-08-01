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

func SafeAtoi(str string) int {
	integer, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return integer
}

func GetProblemSize(buffer *bufio.Reader) int {
	problemSizeStr, _ := buffer.ReadString('\n')
	return SafeAtoi(problemSizeStr)
}

func GetPhoneNumbers(buffer *bufio.Reader) map[string]int {
	problemSize := GetProblemSize(buffer)
	phoneNumbers := make(map[string]int)
	for i := 0; i < problemSize; i++ {
		entry, _ := buffer.ReadString('\n')
		tokens := strings.Split(entry, " ")
		name, phoneNumberStr := tokens[0], tokens[1]
		phoneNumbers[name] = SafeAtoi(phoneNumberStr)
	}
	return phoneNumbers
}

func PrintPhoneNumber(phoneNumbers map[string]int, name string) {
	phoneNumber, ok := phoneNumbers[name]
	if ok {
		fmt.Printf("%v=%v\n", name, phoneNumber)
	} else {
		fmt.Println("Not found")
	}
}

func main() {
	buffer := bufio.NewReader(os.Stdin)
	phoneNumbers := GetPhoneNumbers(buffer)

	for {
		name, err := buffer.ReadString('\n')
		PrintPhoneNumber(phoneNumbers, strings.TrimSpace(name))
		if err != nil {
			// EOF
			break
		}
	}
}
