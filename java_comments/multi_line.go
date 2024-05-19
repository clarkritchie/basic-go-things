package main

import (
	"bufio"
	"fmt"
	"strings"
)

var s2 string = `
foo
bar
/* foo */
bar1

hello
/*
bar2
foo3
*/

moo
// hello
// world
cow
`

func multi_line() {
	scanner := bufio.NewScanner(strings.NewReader(s2))
	scanner.Split(bufio.ScanWords)
	isComment := false
	cleanString := ""
	for scanner.Scan() {
		word := scanner.Text()
		if word == "/*" || word == "//*" {
			isComment = true
		}
		if !isComment {
			cleanString += " " + word
		}
		if word == "*/" {
			isComment = false
		}
	}

	fmt.Printf("original:\n%v\n", s2)
	fmt.Printf("final:\n%v\n", cleanString)
}
