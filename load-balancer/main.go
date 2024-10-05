package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

//bunu biraz daha gelistirip http tabanli lb yapabilirsiniz.

var (
	counter int

	//TODO configurable
	listenAddr = "localhost:8080"

	//TODO configurable
	server = []string{
		"localhost:5001", "localhost:5002", "localhost:5003",
	}
)

func main() {
	//gelen baglantilari kabul etmemiz gerekiyor.
	//tcp dinle.
	//lb in dinleyecegi bir adres ve port belirledik.localhost 8080 de dinlemek istiyorum.
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen %s", err)
	}
	defer listener.Close()

	//process i acik tutmak icin. baglantinin bitmesini beklemek icin.
	for {
		//disardan gelen baglanti
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err)
		}

		backend := chooseBackend()
		fmt.Println("counter=%d backend=%s\n", counter, backend)
		//bu backend e bu baglantiyi proxy et
		//baglantiyi kabul ettikten sonra be e baglanana kadar metod bloke ediyor.
		//yani yeni baglantiyi kabul etmek icin hazir olmuyoruz.
		go func() {
			err := proxy(backend, conn)
			if err != nil {
				log.Printf("WARNIING: proxying failed: %v", err)
			}
		}()
	}

}

func proxy(backend string, c net.Conn) error {
	//tcp üzerinden backend e baglanicam. arka tarafa baglanti veriyor.
	bc, err := net.Dial("tcp", backend)
	if err != nil {
		return fmt.Errorf("failed to connect to backend %s: %v", backend, err)
	}
	//c -> bc
	//arka tarafa bir go routine aciyor. thread e benzeyen concurrent bir islem aslinda.
	go io.Copy(bc, c)
	//bc -> c
	go io.Copy(c, bc)

	return nil
}

func chooseBackend() string {
	//round robin
	s := server[counter%len(server)]
	counter++
	return s
}
