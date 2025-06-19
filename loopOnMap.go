package main

import "fmt"
func main() {
map1 := make(map[string]string)
map1["name"]="darshan"
map1["nature"]="rude"
map1["name2"]="darsh"


for i,v:=range map1{
	fmt.Println(i,v)
}
//just give all the value
for _,v:=range map1{
	fmt.Println(v)
}
//just give all the key
for i,_:=range map1{
	fmt.Println(i)
}

}