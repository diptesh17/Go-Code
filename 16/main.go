package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Strings in GO")

	// 1 : Split
	data := "Diptesh,Anil,Deore"
	parts := strings.Split(data, ",")
	fmt.Println("Split string : ", parts)

	// 2 : Count
	str := "one two one three"
	count := strings.Count(str, "one")
	fmt.Println("One occures times ", count, " Times")

	// 3 : TrimSpace ( only remove pre & post spaces before any character occur )
	str1 := "      I am Dip    "
	trimmed := strings.TrimSpace(str1)
	fmt.Println("Post trim : ", trimmed)

	// 4 : Join ( concate array of strings with a delimeter)
	str11 := "John"
	str12 := "Singh"
	result := strings.Join([]string{str11, str12}, " ")
	fmt.Println("String post join : ", result)
}
