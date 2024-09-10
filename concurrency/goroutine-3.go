package main

import (
	"fmt"
	"time"
)

func main() {
	go boring("boring!")
	fmt.Println("I am listening")
	time.Sleep(time.Second * 2)
	fmt.Println("You are boring. I am leaving.")
}

func boring(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}
