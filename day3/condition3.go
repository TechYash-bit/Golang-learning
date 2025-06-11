//package yash
package main

import "fmt"

//in these program if else-if conditional statement

func main(){

	var age int 
	fmt.Print("enter your age :")
	fmt.Scan(&age)

	if(age>50){
		fmt.Println("you are old man")
	}else if age>=18{
		fmt.Println("you are adult ")
	}else{
		fmt.Println("you are kid")
	}



}