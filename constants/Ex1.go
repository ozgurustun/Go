package main

import (
	"fmt"
	"math"
)

const a int16 = 55

func main() {

	const x float64 = 555.55
	//constants have to be assignable at compile time
	//const errorConstant float64 = math.Sin(1.55)
	const constantOne string = "Constant One"

	//inner constants shadow outer constants
	const a int = 11
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", constantOne, constantOne)
	fmt.Println(math.Sin(x))
}
