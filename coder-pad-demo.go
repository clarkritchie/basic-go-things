package main

import "fmt"

type Lot struct {
	TotalSpaces int
	Remaining   int
}

func (l Lot) RemainingSpaces() int {
	return l.Remaining
}

func main() {
	l := Lot{}
	l.TotalSpaces = 10
	l.Remaining = 5
	r := l.RemainingSpaces()
	fmt.Println("the end: %v\n", r)

	numbers := []int{1, 2, 3, 4}
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println("Sum:", sum)

	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)
}
