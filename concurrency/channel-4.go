package main

func main() {
	unbufferedChan := make(chan int)

	//reader
	go func(ch chan int) {
		value := <-ch
		println(value)
	}(unbufferedChan)

	//writer
	go func(ch chan int) {
		ch <- 1
	}(unbufferedChan)

	//time.Sleep(time.Second)
	println("hello channels")

	// output is non-deterministic. Scheduler probably will not have time to schedule goroutines.
	// so we will not see channel value in the output

}
