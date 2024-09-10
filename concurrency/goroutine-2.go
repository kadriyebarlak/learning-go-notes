package main

import (
	"fmt"
	"time"
)

// go routine inside for-loop without val
func main() {
	for i := 0; i < 10; i++ {
		go func(val int) {
			fmt.Println(val)
		}(i) // i is copy value
	}

	time.Sleep(time.Second * 3)
}
