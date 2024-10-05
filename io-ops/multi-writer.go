package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	mFile, _ := os.Create("mtestFile.txt")
	mfile2, _ := os.Create("mTestFile2.txt")
	mWriter := io.MultiWriter(os.Stdout, mFile, mfile2)

	n, err := mWriter.Write([]byte("multi writer example")) //n=20 byte

	mFile.Close()
	mfile2.Close()

	fmt.Println(err)
	fmt.Println(n)

}
