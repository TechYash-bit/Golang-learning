package main

import "fmt"

//Star printing
//*
//**
//***
//****

func main() {
	fmt.Println("enter the no of row")
	var row int
	fmt.Scan(&row)

	for i := 0; i < row; i++ {
		for j := i; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
