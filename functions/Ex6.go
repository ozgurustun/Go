package main

import "fmt"

func main() {
	a := greeter{
		greeting: "Hello",
		name:     "Ozgur",
	}
	a.greet()
	fmt.Println("Struct name is not changed", a.name)
	a.greetPointer()
	fmt.Println("Name will be changed if we pass pointer of struct", a.name)
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	/*
		Not passing struct,passing copy of struct
	*/
	g.name = "Proof"
}

func (g *greeter) greetPointer() {
	fmt.Println(g.greeting, g.name)
	/*
		Not passing struct,passing copy of struct
	*/
	g.name = "Proof"
}
