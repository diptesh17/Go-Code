package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Data conversion in GO")

	num := 2
	fmt.Printf("Type of num is : %T ", num)

	fmt.Println()

	data := float64(num)
	fmt.Printf("Type of data is : %T ", data)

	// int to string
	no := 123
	str := strconv.Itoa(no)
	fmt.Println("Number : ", no)
	fmt.Println("String : ", str)
	fmt.Printf("Type of string is : %T ", str)

	fmt.Println()

	// string to int : return int & error
	strr := "123"
	to_int, _ := strconv.Atoi(strr)
	fmt.Println("String : ", strr)
	fmt.Println("Number : ", to_int)
	fmt.Printf("Type of to_int is : %T ", to_int)

}
