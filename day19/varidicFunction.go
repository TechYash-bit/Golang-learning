package main

import "fmt"
func main() {

fmt.Println(sum(10,20,30,0))
}

func sum(nums...int)int{
	total:=0
	for _,v:=range nums{
		total+=v
	}
	return total
}