package main

import "fmt"

func integer(num int) {
defer func ()  {
	r:=recover()

	if r!= nil{
		fmt.Print("recover ",r)
	}
}()

	if num < 0 {
		panic("Integer can't be negative")
	}
}

func main() {
	integer(-1)

}