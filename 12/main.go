package main

import "fmt"

func main() {
	fmt.Println("Map in GO")

	// Student map with their marks
	student := make(map[string]int)

	student["Diptesh"] = 85
	student["Ronny"] = 84
	student["Sam"] = 80

	fmt.Println("Marks of Sam : ", student["Sam"])

	// Update :
	student["Sam"] = 70
	fmt.Println("Marks of Sam : ", student["Sam"])

	// Delete :
	delete(student, "Sam")
	fmt.Println("Marks of Sam : ", student["Sam"])

	// Check existance
	mark, exist := student["David"]
	fmt.Println("Marks of David : ", mark)
	fmt.Println("David exist : ", exist)
	marks, exists := student["Diptesh"]
	fmt.Println("Marks of Diptesh : ", marks)
	fmt.Println("Diptesh exist : ", exists)

	// Traverse map using for + range

	for key, val := range student {
		fmt.Printf(" Key is %s and value is %d ", key, val)
	}

	// Initialize at time of declaration
	person := map[string]int{
		"John": 20,
		"Cena": 40,
	}

	for key, val := range person {
		fmt.Printf(" Key is %s and value is %d ", key, val)
	}
}
