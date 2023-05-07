package main

import "fmt"

type idk struct {
	tt int
}

func (p *idk) print_hi() {
	p.tt += 1
	fmt.Println("HI", p.tt)
}

func main() {
	ii := idk{tt: 123}
	ii.print_hi()
	fmt.Println(ii.tt)

	iii := &idk{tt: 123}
	iii.print_hi()
	fmt.Println(iii.tt)
}
