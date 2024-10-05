package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	str1 := strings.NewReader("string reader 1\n")
	str2 := strings.NewReader("string reader 2\n")
	mReader := io.MultiReader(str1, str2)

	io.Copy(os.Stdout, mReader)

}
