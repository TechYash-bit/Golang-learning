package main

import "fmt"
func main() {
//creating a map using make 

map1:=make(map[string]string)
//insert an element in map
map1["name"]="yash"
map1["middle name"]="Niranjan"
map1["surname"]="harne"

//printing a  map

fmt.Println(map1)
//accessing the element in the map
fmt.Println(map1["name"])
fmt.Println(map1["middle name"])
fmt.Println(map1["surname"])

//updating the value of key
map1["name"]="Darshan"
map1["middle name"]="Ashok"
map1["surname"]="Ambekar"
fmt.Println(map1)
//delete an key in map
delete(map1,"middle name")
fmt.Println(map1)
}