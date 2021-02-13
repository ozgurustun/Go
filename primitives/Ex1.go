package main

import "fmt"

func main() {
	b := 1 == 1
	var a bool = true
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T", b, b)
}
