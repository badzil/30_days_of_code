// Solution to https://www.hackerrank.com/challenges/30-2d-arrays.

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

func GetRow(buffer *bufio.Reader) []int {
	rowStr, _ := buffer.ReadString('\n')
	cells := strings.Split(rowStr, " ")
	row := make([]int, len(cells))
	for index, cell := range cells {
		row[index] = SafeAtoi(cell)
	}
	return row
}

func GetGrid() [][]int {
	/* Note that this function is limited to returning 6 rows. */
	buffer := bufio.NewReader(os.Stdin)
	grid := make([][]int, 6)
	for rowIdx := 0; rowIdx < 6; rowIdx++ {
		grid[rowIdx] = GetRow(buffer)
	}
	return grid
}

func GetHourglass(grid [][]int, row, column int) int {
	return grid[row][column] +
		grid[row][column+1] +
		grid[row][column+2] +
		grid[row+1][column+1] +
		grid[row+2][column] +
		grid[row+2][column+1] +
		grid[row+2][column+2]
}

func GetLargestHourglass(grid [][]int) int {
	var largestHourglass int
	for row := 0; row < 4; row++ {
		for column := 0; column < 4; column++ {
			hourglass := GetHourglass(grid, row, column)
			if hourglass > largestHourglass || (row == 0 && column == 0) {
				largestHourglass = hourglass
			}
		}
	}
	return largestHourglass
}

func main() {
	grid := GetGrid()
	fmt.Println(GetLargestHourglass(grid))
}
