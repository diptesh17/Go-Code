package main

import "fmt"

func simple() {
	fmt.Println("Inside Simple Function")
}

// All i/p same : a,b,c int else a int , b string..
func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Inside Main Function")
	simple()

	ans := add(3, 3)
	fmt.Println("Sum of two no's : ", ans)
}
