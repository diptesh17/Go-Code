package main

import "fmt"

func main() {
	fmt.Println("Printf vs Println")

	age := 22
	name := "Diptesh"
	height := 5.5

	// Auto line insert + added spaces post parameters
	fmt.Println("Name :", name, "Age :", age, "Height :", height)
	fmt.Println("Hello")

	// Printf : uses format specifier : "age is %d\n" , age  
}
