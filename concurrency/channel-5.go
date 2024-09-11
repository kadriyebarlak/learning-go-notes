package main

//this creates a buffered channel with a capacity of 1.
//A buffered channel can hold values without requiring a receiver.
//no blocking because the value is stored in the buffer.
func main() {
	bufferedChan := make(chan int, 1)

	bufferedChan <- 1
	bufferedChan <- 2 // buffereddaki value lari okuyup bir sonraki yazma isleminin blockingini kaldiracak bir baska goroutine yok

	println(<-bufferedChan)
}

/*
	channel en fazla 1 değer tutabilir, ve bir sonraki gönderim için önce bu kanaldaki değerin okunması (yani tüketilmesi) gerekir.
	 buffered channel dolduğunda, sender goroutine, channel in bir alıcı tarafından boşaltılmasını bekler.
	 Ancak bu kodda, şu anda channeldaki değeri (1) okuyacak bir alıcı yok, bu yüzden ikinci gönderim bloğa girer ve kilitlenir.
*/
