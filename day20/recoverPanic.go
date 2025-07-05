package main

import "fmt"
func main() {

recovered()
}
func recovered(){
	defer func(){
		if r:=recover();r!=nil{
			fmt.Println("recover ",r)
		}
	}()
	fmt.Println("start process")
	panic("something wrong happen")
}