package main
import "fmt"

func main(){
	fmt.Println("Sclice")

	//declare sclice method 1: uninitialized sclice
	var nums []int

	//length function
	fmt.Println(len(nums))

	//declare sclice method 2: initialized with zero
	//syntax: var Sclice_name=make([]datatype,initial_size,capacity)
	var num1=make([]int,2)
	fmt.Println(num1)

	var num2=make([]int,2,5)
	fmt.Println(num2)

	//creating sclice method 3:
	//empty sclice
	num3:=[]int{}
	fmt.Println(num3)

	//while creating insert an element
	num4:=[]int{1,2,3}
	fmt.Println(num4)

	//capicity function
	fmt.Println("nums : ",cap(nums))
	fmt.Println("num1 : ",cap(num1))
	fmt.Println("num2 : ",cap(num2))
	fmt.Println("num3 : ",cap(num3))
	fmt.Println("num4 : ",cap(num4))
}