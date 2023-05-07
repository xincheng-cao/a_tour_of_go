package main

import (
	"fmt"
	"time"
)

type myDuration int64

func (d myDuration) ToSeconds() float64 {
	return float64(d) / float64(time.Second)
}

func main() {
	//var d Duration = 5 * time.Second
	var d myDuration = myDuration(5 * time.Second)
	k := int64(5 * time.Second)
	fmt.Println(k)
	fmt.Println(d.ToSeconds()) // Output: 5
}
