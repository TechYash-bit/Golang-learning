package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "hello world!"
	str1 := "    1st program     "

	//concatinate an string
	ans := str + " " + str1
	fmt.Println(ans)

	//is string contain something
	fmt.Println(strings.Contains(str, "ll"))

	//get character at index
	fmt.Println(str[6])

	//get substring from index to index
	fmt.Println(str[1:5])

	//conver integer to string
	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println(str3)

	//split an string from some thing
	fmt.Println(strings.Split(str, " "))

	//make sclice to string
	sclice := []string{"india ", "is", "my", "country"}
	fmt.Println(strings.Join(sclice, " "))

	//replace string with some thing
	fmt.Println(strings.Replace(str, "world", "go", -1))

	//remove if having some space at  first
	fmt.Println(str1)
	fmt.Println(strings.TrimPrefix(str, "hello"))
	fmt.Println(strings.TrimSpace(str1) + " a")

	//lowercase and uppercase
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))

	//count some thing in string
	fmt.Println(strings.Count(str, "l"))

	//repate string some time
	fmt.Println(strings.Repeat(str,3))

	//prefix and sufix
	fmt.Println(strings.HasPrefix(str,"hello"))
	fmt.Println(strings.HasPrefix(str,"llo"))
	fmt.Println(strings.HasSuffix(str,"world!"))
	fmt.Println(strings.HasSuffix(str,"world"))

	//utf8 package 
	fmt.Println(utf8.RuneCountInString(str))
	

	//buider package 
	var build strings.Builder

	build.WriteString("hello")
	build.WriteString(" ")
	build.WriteString("world")
	fmt.Println("this will give you an adress ",build)
	fmt.Println(&build)

	build.WriteRune('!')
	fmt.Println(&build)

	build.Reset()
	fmt.Println(&build)

}
