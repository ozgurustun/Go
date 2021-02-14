package main

import "fmt"

func main() {

	aDoctor := struct{ name string }{name: "Ozgur Ustun"}
	anotherDoctor := aDoctor
	//anotherDoctor := &aDoctor
	anotherDoctor.name = "Burak Koken"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)

}
