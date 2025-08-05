package main

import "fmt"

type Stack[T any] struct {
	stack []T
}

func (s *Stack[T]) push(val T) {
	s.stack = append(s.stack, val)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.stack) == 0 {
		var zero T
		return zero, false
	}
	top := s.stack[len(s.stack)-1]
	return top, true
}

func (s Stack[T]) display() {
	if len(s.stack) == 0 {
		fmt.Println("[ ]")
	}
	fmt.Print("[ ")
	for _,v:=range s.stack{
		fmt.Print(v,"  " )
	}
	fmt.Print(" ]")
	fmt.Println()
}

func main(){
	Stack:=Stack[string]{}
	Stack.push("Yash")
	Stack.push("harne")
	Stack.display()
	Stack.pop()
	Stack.display()
}