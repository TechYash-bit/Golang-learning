package main

import (
	"errors"
	"fmt"
)

func main() {
	err := somthingWrong()
	fmt.Println("Error :", err)

}

func somthingWrong() error {
	err := errors.New("somthing wrong happen")
	return err
}
