package main

import (
	"fmt"
	"os"
)
func main() {
fmt.Println("before exit in main function")
//proces()
os.Exit(1)
fmt.Println("after exit in main function")

}
func proces(){
	defer fmt.Println("defer in proces function")
	fmt.Println("before exit in process")
	os.Exit(2)
	fmt.Println("after exit in process function")

}