package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Type conversions

	//Convert int to float32
	var a int = 32

	var b float32
	b = float32(a)
	fmt.Printf("%v,%T", a, a)
	fmt.Printf("\n%v,%T", b, b)

	//Convert int to string
	var c string
	//String conversion package
	c = strconv.Itoa(a)
	fmt.Printf("\n%v, %T", c, c)

}
