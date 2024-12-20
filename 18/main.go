package main

import "fmt"

func main() {

	// defer : executes post main function execution is over
	// Multiple defer : internally uses stack kind of approach 
	fmt.Println("Start")
	defer fmt.Println("Middle")
	fmt.Println("End")
}
