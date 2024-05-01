package main

import (
	"fmt"
	"sort"
)

var arr = []int{1, 2, 3, 4}

func searchArray() bool {
	target := 8
	found := false
	for _, value := range arr {
		if value == target {
			found = true
			break
		}
	}
	return found
}

func sortArray() {
	sort.Ints(arr)
}

func reverseArray() {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func filterArray() {
	original := []int{1, 2, 3, 4, 5}
	filtered := []int{}
	for _, value := range original {
		if value > 2 {
			filtered = append(filtered, value)
		}
	}
}

func deleteSliceItem() {
	slice := []int{1, 2, 3, 4, 5}
	indexToDelete := 2
	slice = append(slice[:indexToDelete], slice[indexToDelete+1:]...)
	fmt.Printf("new slice: %v\n", slice)
}

func copySlice() {
	sourceSlice := []int{1, 2, 3}
	targetSlice := make([]int, len(sourceSlice))
	copy(targetSlice, sourceSlice)
}

func reverseSlice() {
	slice := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func existsInMap() {
	var map_obj = make(map[string]string)
	if val, isExists := map_obj["foo"]; isExists {
		//do steps needed here
		fmt.Printf("val is %v\n", val)
	}
}

type Person struct {
	name string
	age  int
}

func (p *Person) Aging(y int) {
	p.age = p.age + y
}

func main() {
	s := []string{}
	// var s = make([]string, 6)
	s = append(s, "f")

	fmt.Println(s)

	f := searchArray()
	fmt.Printf("Found: %t", f)
	deleteSliceItem()

	p := Person{name: "Bob", age: 4}
	p.age = 30
	fmt.Printf("the age of %s is %v\n", p.name, p.age)
	p.Aging(2)
	fmt.Printf("the age of %s is %v\n", p.name, p.age)
}
