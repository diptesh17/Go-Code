package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isadult"`
}

func main() {
	fmt.Println("JSON in GO")
	person := Person{Name: "John", Age: 30, IsAdult: true}
	fmt.Println(person)

	// Convert person into JSON Encoding ( Marshalling )
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	// Normal jsonData gives array of strings , we need to convert it into strings
	fmt.Println("Marshal : ", string(jsonData))

	// Decoding ( Unmarshalling )
	var data Person
	json.Unmarshal(jsonData, &data)
	fmt.Println("Unmarshal : ", data)
}
