package main

// Simple example of using channels

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go fooBar(ch)
	msg := <-ch
	fmt.Printf("message:  %v", msg)
}

func fooBar(ch chan<- string) {
	time.Sleep(2 * time.Second)
	ch <- "hello world"
}
