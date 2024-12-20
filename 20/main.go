package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Web Request in GO")

	// res type : http response not json
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	defer res.Body.Close()

	// Read :
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Println("Data : ", string(data))
}
