package main

import (
	"fmt"
	"io"
	"time"
)

//for data sharing without a buffer. The reason for using a pipe is generally for use in goroutines.
func main() {
	pReader, pWriter := io.Pipe()
	done := make(chan struct{})

	go read(pReader, done)
	go write(pWriter)

	<-done
}

func read(reader io.Reader, done chan struct{}) {
	buff := make([]byte, 1024)
	for {
		n, err := reader.Read(buff)
		// Even if it returns an error, it might have read some data.
		if n == 0 {
			if err == io.EOF {
				done <- struct{}{} // like signal
				break
			}
			if err != nil {
				fmt.Println(err)
				done <- struct{}{}
				break
			}
		} else {
			fmt.Println(buff[:n])
			if err == io.EOF || err != nil {
				fmt.Println(err)
				done <- struct{}{}
				break
			}
		}
	}
}

func write(writer *io.PipeWriter) {
	//incoming bulk data
	i := 0
	for {
		if i == 10 {
			writer.Close()
			break
		}
		writer.Write([]byte(string(i)))
		i++
		time.Sleep(time.Millisecond * 100)
	}
}
