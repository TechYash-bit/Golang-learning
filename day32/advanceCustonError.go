package main

import (
	"errors"
	"fmt"
)

func main() {

	err:=errorHandling()
	fmt.Print(err)

}
func errorHandling() error{
	err:=somethingWrong()

	if err!=nil{
		return &cusError{
			code: 302,
			message: "wrong ",
			er: err,
		}
	}
	return nil
}

type cusError struct {
	code    int
	message string
	er      error
}

func (m *cusError) Error() string {
	return fmt.Sprintf("Error : %d,%s,%v",m.code,m.message,m.er)
}

func somethingWrong() error{
	return errors.New("custom error")
}