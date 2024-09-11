package main

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 1

	ch2 := make(chan int, 2)
	ch2 <- 2

	var done bool
	for !done {
		select {
		case val1 := <-ch1:
			println(val1)
		case val2 := <-ch2:
			println(val2)
		default:
			done = true
		}
	}

}
