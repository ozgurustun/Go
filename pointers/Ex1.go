package main

import "fmt"

/*

	*int - a pointer to an integer
	Use the address operator & to get address of variable

	Dereference a pointer by preceding with an asterisk(*)
	complex types(e.g structs) are automatically dereferenced

	To create pointers to objects
		Can use address of operator if value type already exists
		ms := myStruct{foo:42}
		p := &ms or directly &myStruct{foo:42}

	All assignment operations in Go are copy operations
	Slices and maps contain internal pointers, so copies point to same underlying data

*/

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, b)
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a, *b)
	*b = 14
	fmt.Println(a, *b)
}
