package main

import "fmt"
func main() {
process(-1.5)

}
func process(nums float32){
	if(nums < 0){
		panic("enter the value greater than zero")
		//after this we will not execute any thing
	}
	fmt.Println("processing ", nums)
	defer fmt.Println("just write a program")

}