package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Print("Enter a name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Print("Enter an address: ")
	var address string
	fmt.Scanln(&address)

	data := map[string]string{
		"name":    name,
		"address": address,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
