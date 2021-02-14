package main

import "fmt"

func main() {
	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Mehmet"
	students[1] = "Ozgur"
	students[2] = "Burak"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student#1: %v\n", students[0])
	fmt.Printf("Number of students %v\n", len(students))
}
