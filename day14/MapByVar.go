package main

import "fmt"
func main() {
//creating an map by using this syntax  "var mapName map[keyType]valuetype"
//❌ You cannot assign values until it's initialized (e.g. with make())
var map2 map[string]string
fmt.Println(map2)
  
map2["movie"]="ms dhoni untold story"
map2["relese year"]="2016"
fmt.Println(map2)
// this code will give you error 
//this syntsx is use for only creating an empty or nil map
}