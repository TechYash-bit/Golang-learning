package main

import "fmt"
func main() {
 map2:=make(map[string]map[string]string)
 map1 := make(map[string]string)
map1["name"]="darshan"
map1["nature"]="rude"
map1["name2"]="darsh"

map2["user"]=map1
fmt.Println(map2)

	// Declare and initialize nested map
	// map2 := make(map[string]map[string]string)

	// // Create inner map
	// map1 := make(map[string]string)
	// map1["name"] = "darshan"
	// map1["nature"] = "rude"
	// map1["name2"] = "darsh"

	// // Assign inner map to outer map
	// map2["user"] = map1

	// // Print the nested map
	// fmt.Println(map2)
}


