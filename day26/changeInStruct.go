package main

import "fmt"

type student struct {
	name       string
	age        int
	percentage float32
}

func changeInMark(s *student ,addmark float32){
	s.percentage+=addmark
}

func main() {
	s := student{
		name:       "yash harne",
		age:        22,
		percentage: 32,
	}

	fmt.Println(s)
	changeInMark(&s,4)
	fmt.Println(s)

}