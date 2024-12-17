package main

import "fmt"

type Person struct {
	FirstName string
	Lastname  string
	Age       int
}

func main() {
	fmt.Println("Struct in GO")

	// method 1
	var dip Person
	fmt.Println("Dip person : ", dip) // by default : __0 ( def value based on data type )
	dip.FirstName = "Diptesh"
	dip.Lastname = "Deore"
	dip.Age = 22
	fmt.Println(dip)

	// method 2
	person2 := Person{
		FirstName: "Sam",
		Lastname:  "Curran",
		Age:       28,
	}

	fmt.Println("Person 2 : ", person2)

	// 3rd Method : new keyword
	var person3 = new(Person)
	person3.FirstName = "Ronni"
	fmt.Println(" Name : ", person3)
}
