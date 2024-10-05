package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var products []Product
var receivedMessages []Product

func main() {

	rabbitmq := NewRabbitMQ()

	go Consumer()

	http.HandleFunc("/product", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			json.NewEncoder(rw).Encode(receivedMessages)
		} else if r.Method == http.MethodPost {
			var newProduct Product
			bodyBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(rw, "Invalid product data", http.StatusBadRequest)
				return
			}
			json.Unmarshal(bodyBytes, &newProduct)
			products = append(products, newProduct)
			fmt.Printf("New Product Received: %+v\n", newProduct)

			body, _ := json.Marshal(newProduct)
			rabbitmq.Publish(body)
			rw.WriteHeader(http.StatusCreated)
		} else {
			http.Error(rw, "unsupporter http method", http.StatusMethodNotAllowed)
		}
	})

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println(err)
		}
	}()

	<-make(chan struct{})
}
