package main

import "fmt"
//this progran for if else statement

func main(){

	var number int
	fmt.Print("enter the any number : ")
	fmt.Scan(&number)

	if number%2==0{
		fmt.Println("the number ",number,"is even")
	}else {

		//the else is written like this
		fmt.Println("the number ",number,"is odd")
	}

}