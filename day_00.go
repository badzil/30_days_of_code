// Solution to https://www.hackerrank.com/challenges/30-hello-world.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	x, _ := stdinReader.ReadString('\n')
	fmt.Println("Hello, World.")
	fmt.Println(x)
}
