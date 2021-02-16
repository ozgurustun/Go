package main

import "fmt"

func main() {
	var divide func(arg1, arg2 float64) (float64, error)
	divide = func(arg1, arg2 float64) (float64, error) {
		if arg2 == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		}
		return arg1 / arg2, nil
	}
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

}
