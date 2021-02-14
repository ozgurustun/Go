package main

import "fmt"

func main() {

	/*
		We also manipulate a in this example, because references are same

	*/

	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	b := append(a[:2], a[3:]...)
	fmt.Println(b)
	fmt.Println(a)
}
