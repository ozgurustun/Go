package main

import "fmt"

func main() {

	number := 50
	guess := 130

	if guess < 1 || guess > 100 {
		fmt.Println("Guess must be between 1 and 100")
	} else if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}

	} else {
		if guess == number {
			fmt.Println("Equals")
		}

		fmt.Println(number <= guess, number >= guess, number != guess)
	}

}
