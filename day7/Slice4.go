package main

import (
	"fmt"
	"slices"
)

func main(){

	var num1 []int
	num1=append(num1, 10,20,30)

	var num2=make([]int,0,5)
	num2 = append(num2, 10,20,30)

	num3:=[]int{10,20,30,40}

	fmt.Println(slices.Equal(num1,num2))
	fmt.Println(slices.Equal(num1,num3))
	fmt.Println(slices.Equal(num3,num2))

	slices.Clone()

}