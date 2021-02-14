package main

import "fmt"

/*
Backed by array

	a := make([]int, 10) //create slice with capacity and length = 10
	a := make([]int, 10, 100) //slice with length = 10 and capacity = 100

*/

func main() {
	a := []int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
}
