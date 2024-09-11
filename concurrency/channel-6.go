package main

func main() {
	bufferedChan := make(chan int, 6)
	done := make(chan bool)

	go func(ch chan int, done chan bool) {
		for val := range ch { //This loop will run until the channel is closed
			println(val)
		}
		done <- true
	}(bufferedChan, done)

	bufferedChan <- 1
	bufferedChan <- 2
	bufferedChan <- 3
	bufferedChan <- 4
	bufferedChan <- 5
	bufferedChan <- 6
	bufferedChan <- 7
	bufferedChan <- 8
	bufferedChan <- 9

	close(bufferedChan) //// close the channel after sending all values

	<-done // // Now it will unblock once done receives a value

	println("main finished.")
}
