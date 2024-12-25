package main

import (
	"fmt"
	"time"
)

func first() {
	fmt.Println("First")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Post First")

}

func second() {
	fmt.Println("Second")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println(" Post Second")

}

func main() {
	fmt.Println("Goroutine !")
	go first()
	go second()

	time.Sleep(1000 * time.Millisecond)

}
