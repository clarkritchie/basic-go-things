package main

// Checking if the given characters are present in a string

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "The ideal candidate will have excellent software development skills"
	string2 := "You will drive our technical vision, strategy and best practices while shipping high quality code every day"
	// Check for presence Using Contains() method of strings package
	res1 := strings.Contains(string1, "Interview")
	res2 := strings.Contains(string2, "technical")
	// Displaying the result
	fmt.Println("Is 'Interview' present in string1 : ", res1)
	fmt.Println("Is 'technical' present in string2 : ", res2)
}
