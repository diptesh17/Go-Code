package main

import "fmt"

func main() {
	fmt.Println("If-else")

	x := 20

	if x > 1 {
		fmt.Println("Its greater than 1")
	} else {
		fmt.Println("Its smaller than 1")
	}

	fmt.Println("Switch case")

	day := 30

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")

	default:
		fmt.Println("Pata nahi")
	}

}
