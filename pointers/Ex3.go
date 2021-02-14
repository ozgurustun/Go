package main

import "fmt"

func main() {
	var a *exStruct
	fmt.Println(a)
	a = new(exStruct)
	fmt.Println(a)
	fmt.Println(*a)
	//a = &exStruct{name: "Ozgur"}
	/*
		Compiler understands below expressions are equal

		(*a).name = "Ozgur"
		fmt.Println((*a).name)

		a. name = "Ozgur"
		fmt.Println(a.name)
	*/
	a.name = "Ozgur"
	fmt.Println(a.name)

}

type exStruct struct {
	name string
}
