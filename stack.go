package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	data []int
}

func (s *Stack) Push(item int) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("stack is empty")
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item, nil
}

func main() {
	stack := Stack{}
	stack.Push(6)
	stack.Push(8)
	stack.Push(1)
	item, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("item is %v\n", item)
	}
}
