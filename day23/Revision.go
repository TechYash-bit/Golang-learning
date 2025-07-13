package main

import (
	"fmt"
	"slices"
)

func main() {

	var name string
	name = "yash"
	fmt.Println(name)

	for i:=0;i<9;i++{
		fmt.Print(name)
		fmt.Println("this is fpr loop")
	}

	i:=0
	for i<5{
		fmt.Println("this is while loop")
		i++
	}
	//var arr [10]int
	arr:=[10]int{12,23,12,34,45,56,67,89,1134}
	fmt.Println(arr)

	for i,v:=range arr{
		fmt.Println("index ",i," value ",v)
	}

	arr2:=[3][3]int{
		{1,2,3},{3,4,5},{5,6,7},
	}
	for i,v:=range arr2{
		fmt.Println("index ",i," value ",v)
		for j,v1:=range v{
			fmt.Println(i,j,"value ",v1)
		}
	}

	//making a sclise
	 var sclice []string
	 sclice=[]string{
		"yash","niranjan","harne",
	 }

	 for i,v:=range sclice{
		fmt.Println("index ",i," value ",v)
	 }

	 s:=make([]string,len(sclice))
	 copy(s,sclice)

	 for i,val:=range s{
		fmt.Println(i,val)
	 }

	 if slices.Equal(s,sclice){
		fmt.Println("true ")
	 }

	 //craeting an map

	 map1:=make(map[int]string)

	 map1[0]="yash"
	// fmt.Println(map1)

	 value,unknownvalue:=map1[0]
	 fmt.Println(" value",value)
	 fmt.Println(unknownvalue)
}