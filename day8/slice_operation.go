package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("this code is to learn about clone,delete,insert,replace function in slice package")

	var num=make([]int,5,6)
	fmt.Println(num)

	num=slices.Insert(num,0,1,2,3,4,5,6,7)
	fmt.Println("after inserting the element : ",num)

	fmt.Println(len(num))
	num=slices.Replace(num,7,12,8,9,10,11,12)
	fmt.Println("after replace : ",num)

	num=slices.Delete(num,1,5)
	fmt.Println("after delete : ",num)

	num1:=slices.Clone(num)
	fmt.Println("num1 : ",num1)
}