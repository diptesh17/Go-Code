package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time package in GO")

	/*

		Go's official time format :
		- "2006-01-02 15:04:05"
		- 2006 : year
		- 01 : month
		- 02 : day
		- 15 : hour
		- 04 : minute
		- 05 : second

	*/

	currentTime := time.Now()
	fmt.Println("Current Time : ", currentTime)

	// currentTime : Data type => Time

	// Format : dd-mm-yy
	fomartt := currentTime.Format("02-01-2006 , 15:04:05")
	fmt.Println("Format : ", fomartt)
}
