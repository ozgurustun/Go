package main

import (
	"fmt"
	"strconv"
)

/*
Variables can be declared as group
*/

var (
	name         string = "Kobe Bryant"
	team         string = "LA Lakers"
	jerseyNumber int    = 8
)

func main() {
	//Cast int to string with strconv.Itoa(string s)
	fmt.Println("Name: " + name + "\nTeam: " + team + "\nJersey Number: " + strconv.Itoa(jerseyNumber))
}
