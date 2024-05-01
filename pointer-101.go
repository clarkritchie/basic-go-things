package main

import "fmt"

func main() {
	fmt.Println("hello")
	s := "a string"
	fmt.Printf("the address of s is %v\n", &s)

	var i int = 5
	var p *int = &i
	fmt.Printf("the address of i is %v\n", &i)
	fmt.Printf("the value of p is %v\n", *p)
}
