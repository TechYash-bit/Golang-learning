package main
import "fmt"

func main() {
err := &myError{
		code:    404,
		message: "Page not found",
	}
	fmt.Println(err.Error())

}

type myError struct{
	code int 
	message string 
}
func ( m *myError) Error() string{
	return fmt.Sprintf("Error : %d,%s",m.code,m.message) 
}