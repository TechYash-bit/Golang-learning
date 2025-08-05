package main

import "fmt"

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top, true
}

func (s Stack[T]) display() {
	if len(s.elements) == 0 {
		fmt.Println("the stack is empty")
	}
	for _,v:=range s.elements{
		fmt.Print(v)
	}
	fmt.Println()
}
func (s Stack[T]) peek() T{
	top:=s.elements[len(s.elements)-1]
	return top
}

func main() {

	stack :=Stack[int]{}
	stack.push(10)
	stack.display()
	fmt.Println(stack.peek()," this is an top element")
}
