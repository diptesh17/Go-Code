package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	/*
		      - File creation + Context insertion

			fmt.Println("File handling in GO")

			file, err := os.Create("Example.txt")
			if err != nil {
				fmt.Println("Error while file creation : ", err)
				return
			}

			content := "This is the new content added in Example.txt file"

			// io.WriteString return : no of bytes , err
			io.WriteString(file , content)
			fmt.Println("Successfully created a file !")
			// close the resource
			defer file.Close()

	*/

	// Reading context of file
	file, err := os.Open("Example.txt")
	if err != nil {
		fmt.Println("Error in file opening : ", err)
		return
	}
	defer file.Close()

	// create buffer
	buffer := make([]byte, 1024)

	// Read context
	for {
		n, err := file.Read(buffer)

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading file : ", err)
			break

		}

		// Process read : read from buffer 0 till n , convert to string & print
		fmt.Println(string(buffer[:n]))
	}
}
