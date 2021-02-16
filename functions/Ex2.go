package main

import "fmt"

/*
pass by value
*/

func main() {
	arg1 := "asd"
	arg2 := "def"
	func2(&arg1, &arg2)
	fmt.Println(arg2)
}

func func2(arg1, arg2 *string) {
	fmt.Println(*arg1, *arg2)
	*arg2 = "ghj"
	fmt.Println(*arg2)
}
