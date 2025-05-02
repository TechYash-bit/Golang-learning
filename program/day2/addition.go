 package main
//package yash

import "fmt"

func main(){
	fmt.Println("enter the first number")
	var num1 float64
	fmt.Scan(&num1)

	var num2 float64
	fmt.Println("enter the second number")
	fmt.Scan(&num2)


	ans:=num1 +num2
	fmt.Println("sum of ",num1,"and",num2 ,"is ",ans)
}