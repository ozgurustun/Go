package main

import (
	"fmt"
)

/*
Package level variables can be declared here.
But can not use := syntax at this level.
Full decleration syntax allowed here.
*/

var a int = 55

func main() {
	fmt.Printf("%v, %T", a, a)
}
