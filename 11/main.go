package main

import "fmt"

func main() {
	fmt.Println("Loops in GO")

	// Only for loop exist in go

	// No from 0 to 9
	for i := 0; i < 10; i++ {
		fmt.Println("Numbers are : ", i)
	}

	// Infinite loop
	// counter := 0
	// for {
	// 	fmt.Println("Infinite loop")
	// 	counter++
	// }

	// Range keyword : simples looping and return idx & it's value

	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Println("Index  and value  : ", index, value)

	}

}
