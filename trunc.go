package main

import "fmt"

func main() {
	var input float64
	fmt.Printf("Enter a number: ")

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	truncated := int(input)

	fmt.Printf("Truncated number: %.d\n", truncated)
}
