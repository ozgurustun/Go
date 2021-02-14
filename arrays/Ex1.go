package main

import "fmt"

/*
Fixed size
Collection of items with same type
Declaration styles
	a := [3]int{1,2,3}
	a := [..]int{1,2,3}
	var a [3]int
*/

func main() {
	//cars := [3]string{"bmw","mercedes","ferrari"}
	cars := [...]string{"bmw", "mercedes", "ferrari"}
	fmt.Printf("Cars: %v", cars)
}
