package main

import "fmt"

func main() {

	//Bitwise operators
	a := 16 // 1100 , 2^4
	b := 7  // 0111

	fmt.Println(a & b)  // 0100
	fmt.Println(a | b)  // 1111
	fmt.Println(a ^ b)  // 1011
	fmt.Println(a &^ b) // 1000

	//Bitshifting

	fmt.Println(a << 3) // 2^4 * 2^3
	fmt.Println(a >> 3) // 2^4 / 2^3

}
