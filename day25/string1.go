package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	str := "my name is yash"
	fmt.Println(str)
	fmt.Println(len(str))

	fmt.Println("iterate with loop")
	for _,v:=range str{
		//fmt.Print(v)//this will give an asii value as output
		fmt.Printf("%c\n",v)
	}

	//creating an char in go

	var a rune='a'
	fmt.Println("this is rune ascii value ",a)
	fmt.Printf("this will print a actual value that is %c",a)
	fmt.Printf("\n%v this will give ascii value",a)

	fmt.Println("this is counting the rune in string",utf8.RuneCountInString(str))
}