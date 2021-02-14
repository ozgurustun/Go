package main

import "fmt"

func main() {
	/*
		Return order of a map is not guaranteed
	*/
	statePopulations := make(map[string]int)

	statePopulations = map[string]int{
		"California": 21222,
		"Texas":      737292,
		"Florida":    37283742,
		"Ohio":       37371,
	}
	fmt.Println(statePopulations)

	/*
		Add element to a map
	*/
	statePopulations["Georgia"] = 123272
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Ohio"])

	/*
		Delete element from a map
	*/

	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)

	//pop, ok := statePopulations["Ohio"]
	//fmt.Println(pop, ok)

	/*
		for instance statePopulations["asd"] returns zero but we cant know that asd key's value is zero or asd key does not exist in the map,
		to check this we can use comma ok syntax, ok will be false if map does not have pop key
	*/

	fmt.Println(statePopulations["Georgia"])
	pop, ok := statePopulations["Georgia"]
	fmt.Println(pop, ok)

	fmt.Println(statePopulations["Georgia"])
	_, ok = statePopulations["Georgia"]
	fmt.Println(ok)

	fmt.Println(len(statePopulations))
}
