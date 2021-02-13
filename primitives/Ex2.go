package main

import "fmt"

func main() {

	var x uint8 = 8
	a := 20
	b := 10
	fmt.Printf("%v, %T\n", x, x)

	//Not allowed to perform math operations with different types
	//fmt.Println(x + a)

	fmt.Println(int(x) + a)

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a * b)
	fmt.Println(a % b)

}
