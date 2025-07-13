package main
import (
	"fmt"
	"unicode/utf16"
)

func main() {
	str := "a"
	runes := []rune(str)               // convert string to runes
	encoded := utf16.Encode(runes)     // encode runes to UTF-16

	fmt.Println("UTF-16 encoded:", encoded)
}
