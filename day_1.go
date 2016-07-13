package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_input() (int, float64, string) {
	reader := bufio.NewReader(os.Stdin)
	first_input, _ := reader.ReadString('\n')
	second_input, _ := reader.ReadString('\n')
	third_input, _ := reader.ReadString('\n')

	// Need to strip spaces otherwise convert methods fail.
	integer, _ := strconv.Atoi(strings.TrimSpace(first_input))
	double, _ := strconv.ParseFloat(strings.TrimSpace(second_input), 64)

	return integer, double, third_input
}

func main() {
	var first_int, second_int int
	var first_double, second_double float64
	var first_string, second_string string

	first_int = 4
	first_double = 4.0
	first_string = "HackerRank "

	second_int, second_double, second_string = get_input()

	fmt.Println(first_int + second_int)
	fmt.Printf("%.1f\n", first_double+second_double)
	fmt.Println(first_string + second_string)
}
