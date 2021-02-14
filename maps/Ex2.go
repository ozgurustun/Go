package main

import "fmt"

func main() {
	/*
		Slices are invalid key type but arrays are valid
	*/
	//a := map[[]int]string{}

	b := map[[3]int]string{}
	fmt.Println(b)
}
