package main

import "fmt"
func main() {
fmt.Println("calling a function say hello")
sayHello()

fmt.Println("calling sum function")
sum:=sum(3,4)
fmt.Println(sum)

fmt.Println("calling function with say hello with name" )
op:=sayhelloWithName("yash")
fmt.Println(op)

fmt.Println("calling function in function ")
fact:=functionInFunction(5)
fmt.Println(fact(2))

}
func sayHello(){
	fmt.Println("hello world")
}

func sum(a,b int) int{
	return a+b
}

func sayhelloWithName(name string) string{
	return"hello "+name
}


func functionInFunction(factorial int) func(int)int{
	return func(x int)int{
		return x*factorial
	}
}