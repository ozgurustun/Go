package main

import "fmt"

func main() {

	x := 10

	for ; x <= 10; x++ {
		fmt.Println(x)
	}

	for x < 10 {
		fmt.Println(x)
		x++
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
}
