package main

import "fmt"

/*
Fixed size
Collection of items with same type
Declaration styles
	a := [3]int{1,2,3}
	a := [..]int{1,2,3}
	var a [3]int

Arrays in Golang are value types unlike other languages like C, C++, and Java where arrays are reference types.

This means that when you assign an array to a new variable or pass an array to a function, the entire array is copied.
So if you make any changes to this copied array, the original array won’t be affected and will remain unchanged.

Tags can be added to struct file to describe field ` `
*/

func main() {
	//cars := [3]string{"bmw","mercedes","ferrari"}
	cars := [...]string{"bmw", "mercedes", "ferrari"}
	fmt.Printf("Cars: %v", cars)
}
