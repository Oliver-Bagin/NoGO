package main

import (
	"fmt"
)

func printSlice(s []string) {
	for _, word := range s {
		fmt.Print(word + " ")
	}
}

func main() {
	// Original slice of words
	letters := []string{"a", "b", "c", "d", "e"}

	printSlice(letters)

	letters = letters.InsertUnique(letters, "f")
	printSlice(letters)

	letters = letters.InsertUnique(letters, "f")
	printSlice(letters)
}
