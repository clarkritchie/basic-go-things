package main

import "fmt"

type Person struct {
	Age int
}

// pointer receiver
func (m *Person) SetAge() {
	m.Age = 42
}

// value receiver
func (m Person) GetAge() int {
	return m.Age
}

func main() {
	p := Person{Age: 10}

	// Using the method with a pointer receiver
	p.SetAge()
	fmt.Println(p.Age)

	fmt.Println(p.GetAge()) // Outputs: 42
}
