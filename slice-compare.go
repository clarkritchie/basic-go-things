package main

// Compare 2 slices

import (
	"bytes"
	"fmt"
)

func main() {
	sl1 := []byte{'C', 'L', 'A', 'R', 'K'}
	sl2 := []byte{'C', 'A', 'T'}
	res := bytes.Compare(sl1, sl2)
	if res == 0 {
		fmt.Println("Equal Slices")
	} else {
		fmt.Println("Unequal Slices")
	}
}
