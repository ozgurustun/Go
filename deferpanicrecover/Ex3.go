package main

import "fmt"

/*
When defer a function it takes the argument at the time when defer is called, not the time of function executed

*/

func main() {
	a := "start"
	defer fmt.Println(a)
	a = "middle"
	defer fmt.Println(a)
	a = "end"
}
