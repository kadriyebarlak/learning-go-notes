package main

import "fmt"

//channel dediğimiz sey synchronize edilmis bir array yani aslinda locklanan bir array.
//bir array e unsafe pointer olarak point ediyor.
// eszamanli olarak sadece 1 data yazilip okunabilir.
// Buffering removes synchronization. verdiginiz size a kadar blocklaniyor.
func main() {

	unbufferedChan := make(chan int)

	// Pass-by-reference. not pass by value
	//ch ile unbufferedChan in arka planda point ettiği array aynı
	go func(ch <-chan int) {
		// blocks until data arrives
		value := <-ch
		fmt.Println(value)
	}(unbufferedChan)

	unbufferedChan <- 1

}
