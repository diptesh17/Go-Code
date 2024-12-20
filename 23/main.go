package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("CRUD in GO")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error in endpoint ", err)
		return
	}

	// Body
	defer res.Body.Close()

	// Validator
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response : ", res.Status)
		return
	}

	// Retrival
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while conversion ", err)
		return
	}

	fmt.Println("TOdo 1 : ", string(data))

}
