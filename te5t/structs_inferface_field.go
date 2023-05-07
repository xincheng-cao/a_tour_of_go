package main

import "fmt"

type MyInterface interface {
	Method()
}

type MyStruct struct {
	InterfaceField MyInterface
	Lol            int
	Bushi          string
}

func main() {
	var s MyStruct
	fmt.Println(s.InterfaceField, s.Lol, s.Bushi) // Output: <nil>
	d := MyStruct{}
	fmt.Println(d.InterfaceField, d.Lol, d.Bushi)

	var c MyStruct = MyStruct{Lol: 0}
	fmt.Println(c.InterfaceField, c.Lol, c.Bushi)
}
