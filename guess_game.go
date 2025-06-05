package main 

import(
	"fmt"
	"math/rand"
	"time"
)

func main(){

	source:=rand.NewSource(time.Now().UnixNano())
	random:=rand.New(source)
	target:=random.Intn(100)+1

	fmt.Println("Welcome to this game")
	fmt.Println("I have chosen a number between 1 to 100 ")
	fmt.Println("can you guess this number")
	var guess int


	for{
		fmt.Println("Enter your guess")
		fmt.Scanln(&guess)

		if guess==target{
			fmt.Println("you guess correct number")
			break
		}else if guess>target{
			fmt.Println("you predict big number .target number is very low")
		}else{
			fmt.Println("you predict small number .target is such a big number")
		}
	}
	}
