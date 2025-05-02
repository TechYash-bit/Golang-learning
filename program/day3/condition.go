package main

import "fmt"
//how to use conditional statement if 

func main(){

	var age int 
	fmt.Println("enter the age")

	fmt.Scan(&age)

	if age>=18 {
		fmt.Println("your are adult")
	}

}