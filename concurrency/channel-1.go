package main

import "fmt"

// channel as a first-class value.
// in Go, you dont talk to a process
// you talk to a channel and the other end of that channel is some other thing that could be reading the values.
// the cahnnel idea is more like writing to a file descriptor rather than writing to a file by name.
// A channel provides a connection between two goroutines.
// Synchronization : A sender and receiver must both be ready to play their part in the communication. Otherwise we wait until they are.
// Do not communicate by sharing memory, share memory by communication:
// you do not have some blob of memory and then put locks and mutexes instead you use the channel to pass the data back and forth.
func main() {

	c := make(chan int)

	//sending is waiting for a receiver, but no receiver is started. deadlock occurs.
	c <- 1

	val := <-c

	fmt.Println(val)

}
