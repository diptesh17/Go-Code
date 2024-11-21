package main

import "fmt"

// If want to return string
// func divide(a, b float64) (float64, string) {

// 	if b == 0 {
// 		return 0, "Deno is 0"
// 	}

// 	return a / b, "nil"
// }

func divide(a, b float64) (float64, error) {

	if b == 0 {
		return 0, fmt.Errorf("deno is 0")
	}

	return a / b, nil
}

func main() {
	fmt.Println("Error Handling :) ")

	ans, err := divide(10, 0)

	if err == nil {
		fmt.Println(ans)
	}

	fmt.Println(err)

}
