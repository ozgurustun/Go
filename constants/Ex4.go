package main

import "fmt"

const (
	errorSpecialist = iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main() {
	var specialistType int
	fmt.Printf("%v\n", specialistType == catSpecialist)
	fmt.Printf("%v\n", specialistType == errorSpecialist)
}
