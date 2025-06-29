package main
import "fmt"

func main(){

process()
}
func process(){
	fmt.Println("this is no defer satement")
	defer fmt.Println("this is defer statement")
	fmt.Println("this is the statement after defer")
}