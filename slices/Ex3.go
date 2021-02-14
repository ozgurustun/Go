package main

import "fmt"

func main() {

	//slices dont have to have a size like array, we can add and remove elements from slices

	//a := make([]int, 3)
	a := make([]int, 3, 100)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
}
