package main

import (
	"fmt"
	"strconv"
)

func main() {
	var table [5][6]string
	for row := 0; row < 5; row++ {
		for column := 0; column < 6; column++ {
			table[row][column] =
				strconv.Itoa(row) + "," +
					strconv.Itoa(column)
		}
	}
	fmt.Println(table)
}
