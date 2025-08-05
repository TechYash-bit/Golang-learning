package main

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	empId int
	person
	salary     float64
	department string
}

func main() {
	e := employee{
		empId: 10,
		person: person{
			name: "yash",
			age:  22,
		},
		salary:     12345.34,
		department: "devlopment",
	}

	fmt.Println(e.name)

}
