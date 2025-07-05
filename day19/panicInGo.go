package main

import "fmt"
func main() {

proces(1)
proces(-10)
}

func proces(nums int){
	defer fmt.Println("statement 1")
	defer fmt.Println("statement 2")

	if nums<0{
		panic("enter ther value greter then zero")
		//after this we will not execute any thing
	}
	fmt.Println("processing ",nums)
}