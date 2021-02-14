package main

import "fmt"

func main() {
	statePopulations := map[string]int{
		"California": 21222,
		"Texas":      737292,
		"Florida":    37283742,
		"Ohio":       37371,
	}

	fmt.Println(statePopulations)

	/*

		Side effect

	*/

	s := statePopulations
	delete(s, "Texas")

	fmt.Println(s)
	fmt.Println(statePopulations)
}
