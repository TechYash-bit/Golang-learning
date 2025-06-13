package main

import "fmt"
func main() {

//creating a map by using an "mapvarible = map[keytype]valuetype"
//creting a map for storing mark of student

map1:=map[string]int{
	"science":23,
	"math":24,
	"english":23,

}
fmt.Println(map1)
//clear the map
clear(map1)
fmt.Println(map1)
}