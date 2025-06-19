package main

import "fmt"
func main() {
	// map1:=make(map[string]map[string]string)
	// map2:=make(map[string]string)
	// map2["name"]="yash"
	// map2["password"]="7418@dar"

	// map1["user id"]=map2
	// fmt.Print(map1)

	// for i,v:=range map1{
	// 	fmt.Println(i,v)
	// 	for i,map2:=range map2{
	// 		//fmt.Println("en")
	// 		fmt.Println(i,map2)
	// 	}
	// }

	map3:=make(map[string]map[string]string)
	map3["user1"]=make(map[string]string)
	map3["user2"]=make(map[string]string)
	map3["user3"]=make(map[string]string)
	map3["user4"]=make(map[string]string)
	map3["user1"]["name"]="yash"
	map3["user2"]["name"]="rohan"
	map3["user3"]["name"]="om"
	map3["user4"]["name"]="rahul"
	map3["user1"]["paswd"]="harne"
	map3["user2"]["paswd"]="uke"
	map3["user3"]["paswd"]="kale"
	map3["user4"]["paswd"]="thakare"

	for i,user:=range map3{
		fmt.Println(i)
		for i,v:=range user{
		fmt.Println(i,v)
		}
	}

	
	
}


