package main

// Simple example of using ch1annels

import (
	"fmt"
	"time"
)

func bar(ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("bar %v", i)
	}
	close(ch)
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go foo(ch1)
	go bar(ch2)
	go bar(ch1)

	msg1 := <-ch1
	fmt.Printf("message 1: %v\n", msg1)

	for msg2 := range ch2 {
		fmt.Printf("message 2: %v\n", msg2)
	}
}

func foo(ch chan<- string) {
	time.Sleep(2 * time.Second)
	ch <- "foo"
}
