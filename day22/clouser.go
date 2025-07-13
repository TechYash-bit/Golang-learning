package main

import "fmt"
func main() {
value:=adder();
fmt.Println(value)//this will print the adderes of function
fmt.Println(value())
fmt.Println(value())
fmt.Println(value())
fmt.Println(value())
multiplerValue:=multipler()
fmt.Println(multiplerValue(3))
fmt.Println(multiplerValue(4))
fmt.Println(multiplerValue(5))

}

func adder()func()int{
	i:=0
	return func()int{
		i++
		return i
	}
}
func multipler()func(x int)int{
	i:=2
	return func(x int)int{
		i=i*x
		return i
	}
}