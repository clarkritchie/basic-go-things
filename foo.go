package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Number int    `json:"someNumber"` // <----  this is called a "json hint"
	Title  string `json:"aTitle"`
}

func main() {
	foo_marshalled, _ := json.Marshal(Foo{Number: 1, Title: "test"})
	fmt.Printf("%v\n", string(foo_marshalled))
}
