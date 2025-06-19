package main

import (
	"fmt"
	"maps"
)
func main() {

	mp:=make(map[string]string)
	mp["name"]="yash"
	mp["surname"]="harne"
	fmt.Println("mp map print ")
	fmt.Println(mp)

	mp2:=make(map[string]string)
	mp2["name"]="yash"
	mp2["surname"]="harne"
		fmt.Println("map2 ")
	fmt.Println(mp2)
	//check whethere the map is same or not
	if maps.Equal(mp,mp2){
		fmt.Println("the two map are equal")
	}

}