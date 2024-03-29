package main

import (
	"flag"
	"fmt"
	"strings"
)

const SEPARATOR = "   "

func rowLength(row []int64) int {
	length := 0
	for _, num := range row {
		length += len(fmt.Sprintf("%d", num)) + len(SEPARATOR)
	}
	length -= len(SEPARATOR)

	return length
}

func pascalTriangle(rows *int64, color *bool) {
	var rowsData = make([][]int64, *rows)
	var row, col int64
	var number int64

	for row = 0; row < *rows; row++ {
		number = 1
		for col = 0; col <= row; col++ {
			if col > 0 {
				number = number * (row - col + 1) / col
			}
			rowsData[row] = append(rowsData[row], number)
		}
	}

	maxWidth := rowLength(rowsData[*rows-1])

	for _, r := range rowsData {
		var rowOutput = ""

		var padding = (maxWidth - rowLength(r)) / 2
		fmt.Print(strings.Repeat(" ", padding))

		for _, num := range r {
			rowOutput += fmt.Sprint(num, SEPARATOR)
		}

		if *color {
			fmt.Printf("\033[1;34m%s\033[0m", rowOutput)
		} else {
			fmt.Printf("%s", rowOutput)
		}

		fmt.Println()
	}

}

func main() {
	var rows = flag.Int64("rows", 5, "Number of rows in Pascal's Triangle")
	var color = flag.Bool("color", false, "Should we color the output?")
	flag.Parse()

	pascalTriangle(rows, color)
}
