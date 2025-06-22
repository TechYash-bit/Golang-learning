package main

import "fmt"
func main() {
sclice:=make([]int,5)

result:=append([]int{2},sclice...)
fmt.Println(result)
}