package main

import (
	"fmt"
	"sync"
)

// concurrency is the composition(a whole) of independently executing computations.
// concurrency is a way to structure software that simulate the real world.
// go routine is an independently execution function. goroutine says run the function but i do not wait the answer.
// it has own call stack unlike some threading libraries. how big the stack is. it is never an issue in Go.
// It is not a thread. there might be only one thread with thousands of goroutines.
// goroutines are multiplexed dynamically onto threads as needed to keep all the goroutines running.
func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		fmt.Println("Hello from goroutine!")
		wg.Done()
	}()

	//blocking
	wg.Wait()

	fmt.Println("Hello from main!")
}
