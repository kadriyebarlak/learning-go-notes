package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//readFile()
	//readFileBufio()
	readFileSeek()
}

func readFile() {
	r, err := os.ReadFile("testFileWrite.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(r))
	}
}

func readFileBufio() {
	r, err := os.Open("testFileWrite.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		bufReader := bufio.NewReader(r)
		io.Copy(os.Stdout, bufReader)
	}
}

func readFileSeek() {
	f, _ := os.Open("testFile.txt")
	f.Seek(5, 1)
	readByte := make([]byte, 4)
	f.Read(readByte)

	fmt.Println(string(readByte))
}
