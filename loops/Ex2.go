package main

import "fmt"

func main() {

	/*
		Looping over collection(arrays,slices,maps,strings,channels)
	*/

	s := []int{2, 4, 6}
	//Get indexes and values
	for k, v := range s {
		fmt.Println(k, v)
	}

	a := map[string]int{
		"California": 1234,
		"Texas":      4312,
	}

	//get keys and values
	for k, v := range a {
		fmt.Println(k, v)
	}

	//get only keys
	for k := range a {
		fmt.Println(k)
	}

	//get only values
	for _, v := range a {
		fmt.Println(v)
	}

	b := "hello"
	for k, v := range b {
		fmt.Println(k, string(v))
	}
}
