package main

import "fmt"

func main() {
    fmt.Print("Enter the number: ")
    var num int
    fmt.Scan(&num)

    if num < 0 {
        fmt.Println("Enter a number greater than or equal to 0. Factorial of negative numbers doesn't exist.")
        return
    }

    fact := 1
    for i := 1; i <= num; i++ {
        fact *= i
    }

    fmt.Printf("Factorial of %d = %d\n", num, fact)
}
