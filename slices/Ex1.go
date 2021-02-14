package main

import "fmt"

/*
Reference type
Backed by array

	a := make([]int, 10) //create slice with capacity and length = 10
	a := make([]int, 10, 100) //slice with length = 10 and capacity = 100

So when you pass a slice to a function, a copy will be made from this header, including the pointer, which will point to the same backing array.
Modifying the elements of the slice implies modifying the elements of the backing array, and so all slices which share the same backing array will "observe" the change.
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
