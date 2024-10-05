package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3)

	for {
		var input string
		fmt.Printf("Enter an integer or 'X' to quit: ")

		fmt.Scanln(&input)

		if input == "X" || input == "x" {
			fmt.Println("Existing.")
			break
		}
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input.")
			continue
		}
		slice = append(slice, num)
		sort.Ints(slice)

		fmt.Println("Sorted slice: ", slice)
	}
}
