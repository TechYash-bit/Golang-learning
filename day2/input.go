package main

import "fmt"

func main(){
	//enter the input type int
	var i int
	fmt.Print("Type a number: ")
  fmt.Scan(&i)
  fmt.Println("Your number is:", i)

  var name string
  fmt.Println("enter your first name ")
  fmt.Scan(&name)
  fmt.Println("your name is ",name)


  var price float32
  fmt.Println("enter the price of product")
  fmt.Scan(&price)
  fmt.Println("the price of product is ",price)

}