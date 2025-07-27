package main

import "fmt"

type introduction struct{
	name string 
	age int
}
type employee struct{
	emp_id int
	department string
	name string
	salary float32
	position string 
}


func printStruct(e employee) {
fmt.Println("Employee id: ",e.emp_id)
fmt.Println("Department : ",e.department)
fmt.Println("Name : ",e.name)
fmt.Println("Position : ",e.position)
fmt.Println("Salary : ",e.salary)
}
	



func main(){
p:=introduction{
	name:"yash",
	age :22,
}

fmt.Printf("my name is %s and age is %d \n",p.name,p.age)

e:=employee{
	name:"yash harne",
	emp_id: 23,
	position: "devloper",
	salary: 34567,
	department: "software",
}
printStruct(e)
}