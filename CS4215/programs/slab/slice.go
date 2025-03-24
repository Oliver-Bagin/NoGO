package main

import "fmt"

func main() {
	// Create a slice of strings
	greetings := []string{"Hello", "world", "from", "Go", "slices!"}

	// Loop through the slice and print each word
	for _, word := range greetings {
		fmt.Print(word + " ")
	}

	fmt.Println() // Add a newline at the end
}
