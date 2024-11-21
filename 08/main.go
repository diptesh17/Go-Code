package main

import "fmt"

func main() {
	fmt.Println("Arrays in GO ")

	var name [5]string
	name[0] = "Diptesh"
	name[1] = "Deore"
	fmt.Println("Your name : ", name)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
}
