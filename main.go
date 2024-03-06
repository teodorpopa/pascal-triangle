package main

import "flag"

func pascalTriangle(rows *int64) {
	var row, col int64
	var number int64

	for row = 0; row < *rows; row++ {
		number = 1
		for col = 0; col <= row; col++ {
			if col > 0 {
				number = number * (row - col + 1) / col
			}
			print(number, " ")
		}
		println()
	}
}

func main() {
	var rows = flag.Int64("rows", 5, "Number of rows in Pascal's Triangle")

	pascalTriangle(rows)
}
