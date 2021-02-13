package main

import "fmt"

//const(
//	a = iota
//	b = iota
//	c = iota
//)

const (
	a = iota
	b
	c
)

const (
	d = iota
)

func main() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
}
