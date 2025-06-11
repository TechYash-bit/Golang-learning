package main

import "fmt"

func main() {
	fmt.Println("Question :- find the largest element in array")

	
	var arr [10]int

	for i:=0;i<len(arr);i++ {
		fmt.Println("enter the ",i+1,"element  ")
		fmt.Scan(&arr[i])
	}

	large:=0
	index:=-1
	 

	for j:=0;j<len(arr);j++ {
		if arr[j] > large{
			large=arr[j]
			index=j
		}
	}

	fmt.Println("largest element is ",large," at index ",index)


}