package main

import "fmt"
func main(){
	fmt.Println("enter any data type in which you give an input  (int-type/string-type/float-type/bool-type)")

	var datatype string
	fmt.Scan(&datatype)

	switch datatype{
	case "string-type":
		fmt.Println("you give  string type of input")
	case "int-type":
		fmt.Println("you give int type input")
	case "float-type":
		fmt.Println("you give float type input")
	case "bool-type":
		fmt.Println("you give bool type input")
	default:
		fmt.Println("please enter correct data type input")
	}
}