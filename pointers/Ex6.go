package main

import "fmt"

func main() {
	a := 3
	b := &a
	fmt.Println("Variable a memory address : ", &a)
	fmt.Println("Pointer b value: ", b, " Pointer b memory address(pointer): ", &b)
	c := &b
	fmt.Println("Pointer c value(which is b memory address): ", c)
	fmt.Println("Pointer c memory address(pointer): ", &c)
	fmt.Println("Deference once of c pointer value(which is b variable value): ", *c)
	fmt.Println("Deference two times of c pointer value(which is a variable value): ", **c)
}
