package main

import "fmt"

func main() {
	var size int
	fmt.Println("Enter the size of array")
	fmt.Scan(&size)
	var arr = make([]int,size)

	for i:=0;i<len(arr);i++ {
		fmt.Print("Enter the ",i+1," element : ")
		fmt.Scan(&arr[i])
	}

	sum:=0

	for i:=0;i<len(arr);i++ {
		sum+=arr[i]
	}

	avg:=sum/len(arr)

	fmt.Println("sum = ",sum)
	fmt.Println("Average = ",avg)

}