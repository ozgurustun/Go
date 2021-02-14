package main

import "fmt"

func main() {
	//a := [...]int{1,2,3}
	//b := a
	//b[1] = 5
	//fmt.Println(a)
	//fmt.Println(b)

	//b points same data with a

	a := [...]int{1, 2, 3}
	b := &a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}
