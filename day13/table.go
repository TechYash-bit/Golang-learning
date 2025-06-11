package main

import "fmt"
func main() {

	fmt.Println("Enter the number which table you have to print ")
	var a int 
	fmt.Scan(&a)

	for i:=1;i<11;i++{
		fmt.Println(a," * ",i," = ",a*i)
	}

}