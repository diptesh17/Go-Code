package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// fmt.Println("User Input : ")

	// Scan : read till white space
	// var name string
	// fmt.Println("Enter your name : ")
	// fmt.Scan(&name)
	// fmt.Println("Your name is : ", name)

	// For full line : bufio
	fmt.Println("Enter your Full Name : ")
	reader := bufio.NewReader(os.Stdin)
	fullname, _ := reader.ReadString('\n')
	fmt.Println("Full Name : ", fullname)
}
