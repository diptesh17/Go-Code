package main

import "fmt"

func main() {
	fmt.Println("Pointers in GO")

	num := 2
	ptr := &num

	fmt.Println("Number : ", num)
	fmt.Println("Address : ", ptr)
	fmt.Println("Value : ", *ptr)
}
