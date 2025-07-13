package main

import "fmt"

func main() {
fmt.Println("this is main function")
}
func init() {
	fmt.Println("this is init function")// this will run before the main function
}