package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Printf("Enter a string to check for 'i', 'a' and 'n': ")

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if strings.HasPrefix(input, "i") && strings.Contains(input, "a") && strings.HasSuffix(input, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
