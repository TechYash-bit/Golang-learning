package main

import (
	"fmt"
	"reflect"
)

//insertion operation
func main(){
	fmt.Println("Sclice 2")
	//creating sclice
	var num=make([]int,2,5)
	fmt.Println("num ",num)

	//Adding element in  sclice
	num=append(num, 4)
	fmt.Println("after :",num)

	

	//creating sclice with pre inserted element 

	num1:=[]int{1,2,3}
	fmt.Println(num1)
	num1=append(num1, 3)
	fmt.Println("after :",num1)

	//creating empty sclice
	var num3 []int 
	fmt.Println(num3)
	num3=append(num3, 4)
	fmt.Println("after :",num3)


	fmt.Println("type : ",reflect.TypeOf(num))
}