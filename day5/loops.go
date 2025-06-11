package main
//in this program we use for loop as while loop  beacause while loop is not present in go lang

import "fmt"

func main(){
	fmt.Println("enter your name")
	var name string
	fmt.Scan(&name)

	fmt.Println("how many times you print your name ")
	var time int
	fmt.Scan(&time)

	i:=0

	for i<time {
		fmt.Println(i+1,"] ",name)
		i++
	}
}