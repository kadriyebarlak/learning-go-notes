package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	sReader := strings.NewReader("test string")

	//duplicates the stream
	tReader := io.TeeReader(sReader, os.Stdout)

	fmt.Println("starting..")

	//to read all the data from the stream.
	readBytes, _ := io.ReadAll(tReader)

	fmt.Println("\nread string.")

	fmt.Println(string(readBytes))
}
