package main

import "fmt"

/*
embedded concept(instead of inheritance)
*/

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	/*
		embed the enimal struct into the Bird struct
		Bird is not type of Animal it is independet struct
		not is a relationship, has a relationship
	*/
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {

	/*
		b:= Bird{}
			b.Name = "Emu"
			b.Origin = "Australio"
			b.SpeedKPH = 48
			b.CanFly = false
	*/

	b := Bird{
		Animal: Animal{
			Name:   "Emu",
			Origin: "Australio",
		},
		SpeedKPH: 88,
		CanFly:   false,
	}

	fmt.Println(b)
}
