package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	os.WriteFile("lines.txt", []byte("line 1\nline 2\nline 3"), os.ModePerm)

	f, err := os.Open("lines.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}