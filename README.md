# 🚀 My Golang Learning Journey

<div align="center">
  
![Go Gopher](https://github.com/TechYash-bit/Golang-learning/blob/main/program/image/golanglogo.png)

*Documenting my progress learning the Go programming language*
</div>

## 📚 Table of Contents
- [Introduction](#introduction)
- [Code Examples](#code-examples)
  - [Hello World](#hello-world)
  - [Variables and Types](#variables-and-types)
  - [Control Structures](#control-structures)
  - [Functions](#functions)
- [Language Features](#language-features)
- [Learning Resources](#learning-resources)
- [Progress Tracker](#progress-tracker)
- [Setup Instructions](#setup-instructions)

## Introduction
This repository documents my journey learning Go programming language. As I explore Go's simplicity and power, I'm recording code examples, insights, and progress to track my growth and serve as a reference for future projects.

## Code Examples

### Hello World
```go
package main
import "fmt"
func main() {
    fmt.Println("Hello, World!")
}
```

**Output:**
```
Hello, World!
```

![Hello World Output](https://github.com/TechYash-bit/Golang-learning/blob/main/program/image/Day1-hello.png)

---

### Variables and Types
```go
package main

import "fmt"

func main() {

	//creating a string variable
	var name string = "yash niranjan harne"
	fmt.Println(name)

	var name2="yash"
	fmt.Println(name2)

	name3:="harne"
	fmt.Println(name3)

	//creating an int variable
	var age int=12
	fmt.Println(age)

	var age2=123
	fmt.Println(age2)

	age3:=1234

	fmt.Println(age3)
	//craeting an float variable
	var  price float32=12.36
	fmt.Println(price)

	var price2=13.36
	fmt.Println(price2)

	price3:=123.36
	fmt.Println(price3)

}
-----
outpur:-

yash niranjan harne
yash
harne
12
123
1234
12.36
13.36
123.36
```

![Variables Output](https://github.com/TechYash-bit/Golang-learning/blob/main/program/image/Screenshot%20(43).png)



---

## Language Features

<div align="center">

| Feature | Description |
|---------|-------------|
| 🔄 **Concurrency** | Built-in goroutines and channels for efficient concurrent programming |
| ⚡ **Performance** | Compiled language with speed comparable to C/C++ |
| 🧰 **Simplicity** | Clean syntax, garbage collection, and minimal language complexity |
| 📦 **Standard Library** | Comprehensive built-in packages for common tasks |
| 🔄 **Cross-Platform** | Compiles to native binaries for all major operating systems |

</div>

## Learning Resources
- 📚 [A Tour of Go](https://tour.golang.org/) - Official interactive Go tutorial
- 📘 [Go by Example](https://gobyexample.com/) - Hands-on introduction to Go
- 📺 [Gophercises](https://gophercises.com/) - Free coding exercises for budding Gophers
- 📖 [The Go Programming Language](https://www.gopl.io/) - Comprehensive book by Alan A. A. Donovan and Brian W. Kernighan
- 🌐 [Go Documentation](https://golang.org/doc/) - Official documentation

## Progress Tracker

- [x] **Basics**
  - [x] Hello World
  - [x] Variables and Types
 

## Setup Instructions

1. **Install Go**:
   - Download from [golang.org/dl/](https://golang.org/dl/)
   - Verify installation with `go version`

2. **Run Examples**:
   - Save code to a `.go` file
   - Execute with `go run filename.go`




**📝 Last Updated:** April 26, 2025

🔍 **Follow my progress as I continue exploring Go!** 🔍

