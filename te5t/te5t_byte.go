package main

import "fmt"

func main() {
	var ch byte = 'A'
	var ui uint8 = 65
	fmt.Printf("ASCII value of '%c' is %d\n", ch, ch)
	fmt.Printf("ASCII value of '%c' is %d\n", ui, ui)

	ch = 'a'
	fmt.Printf("ASCII value of '%c' is %d\n", ch, ch)
}
