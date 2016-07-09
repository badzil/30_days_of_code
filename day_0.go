// Solution to https://www.hackerrank.com/challenges/30-hello-world.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin_reader := bufio.NewReader(os.Stdin)
	x, _ := stdin_reader.ReadString('\n')
	fmt.Println("Hello, World.")
	fmt.Println(x)
}
