package main

import "fmt"

var i int = 5

func main() {

	//Inner scope is used
	fmt.Println(i)
	var i int = 7
	fmt.Println(i)
	// Declared variables have to be used
	// If you have declared local variables and not used, compile time error
	x := 5
}
