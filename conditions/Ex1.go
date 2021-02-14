package main

import "fmt"

func main() {
	statePopulations := map[string]int{
		"California": 12345,
		"Texas":      12345,
		"Florida":    12345,
		"New York":   12345,
		"Ohio":       12345,
	}

	if pop, ok := statePopulations["Ohio"]; ok {
		fmt.Println(pop)
	}
}
