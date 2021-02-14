package main

import "fmt"

func main() {
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is an int")
	case string:
		fmt.Println("i is an int")
	default:
		fmt.Println("i is another type")
	}
}
