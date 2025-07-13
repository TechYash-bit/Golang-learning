package main

import "fmt"

func main() {
print(4)
revPrint(4)

val:=additionOfN(10)
fmt.Print("value ",val)
}

func print(x int) {

	if x == 0 {
		return
	}
	fmt.Println(x)
	print(x-1)
}

func revPrint(x int){
	if x==0{
		return
	}
	revPrint(x-1)
	fmt.Println(x)
}

func additionOfN(x int)int{
	

	if x==0{
		return 0
	}
	return x+additionOfN(x-1)
	
}