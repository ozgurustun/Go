package main

import "fmt"

func main() {
	//var x float64 = 3.14

	x := 3.14
	x = 13.7e72
	x = 2.1e14
	fmt.Printf("%v, %T", x, x)

}
