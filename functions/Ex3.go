package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(*sumPointer(1, 2, 3, 4, 5))
	fmt.Println(result(1, 2, 3, 4, 5))
	fmt.Println(divide(2.0, 0.0))
	a, err := divide(6.7, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a)
}

func sum(args ...int) int {
	result := 0
	for _, v := range args {
		result += v
	}
	return result
}

func sumPointer(args ...int) *int {
	result := 0
	for _, v := range args {
		result += v
	}
	return &result
}

func result(args ...int) (result int) {
	for _, v := range args {
		result += v
	}
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
