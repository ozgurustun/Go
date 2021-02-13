package main

import "fmt"

func main() {

	//Strings are immutable

	x := "Kobe Bryant"
	fmt.Printf("%v, %T", x, x)
	fmt.Printf("\n%v, %T", x[2], x[2])
	fmt.Printf("\n%v, %T", string(x[2]), x[2])

	//Concat
	a := "Kobe"
	b := "Bryant"
	fmt.Printf("\n%v, %T", a+b, a+b)

}
