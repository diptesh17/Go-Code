package main

import "fmt"

func main() {
	fmt.Println("Slices in Golang :) ")

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Length of Slice Before Append : ", len(numbers))
	fmt.Println("Number : ", numbers)

	numbers = append(numbers, 6, 7, 8, 9, 10)
	fmt.Println("Length of Slice After Append : ", len(numbers))
	fmt.Println("Number : ", numbers)

	// Internals of slices
	fmt.Println("Slice : ", numbers)
	fmt.Println("Length : ", len(numbers))
	fmt.Println("Capacity : ", cap(numbers))

	// Make slice using make function (datatype , len , cap)

	nos := make([]int, 3, 5)
	fmt.Println("Slice : ", nos)
	fmt.Println("Length : ", len(nos))
	fmt.Println("Capacity : ", cap(nos))

}
