package main

import "fmt"

/*
array is a valid key type but slices are not
Check presence with "value,ok" form

Map types are reference types
*/

func main() {
	statePopulations := map[string]int{
		"California": 21222,
		"Texas":      737292,
		"Florida":    37283742,
		"Ohio":       37371,
	}
	fmt.Println(statePopulations)
}
