package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous function")
	}()

	f := func() {
		fmt.Println("Anonymous function")
	}

	f()
}
