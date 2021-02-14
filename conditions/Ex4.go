package main

import "fmt"

func main() {

	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough
	case i >= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}
}
