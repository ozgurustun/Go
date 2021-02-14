package main

import "fmt"

func main() {

	x := 1

	switch x {
	case 0, 4, 8:
		fmt.Println("Zero, four or eight")
		break
	case 1, -1:
		fmt.Println("One or minus one")
		break
	default:
		fmt.Println("another number")
		break
	}

	switch i := 2 + 3; i {
	case 0, 4, 8:
		fmt.Println("Zero, four or eight")
		break
	case 1, -1:
		fmt.Println("One or minus one")
		break
	default:
		fmt.Println("another number")
		break
	}
}
