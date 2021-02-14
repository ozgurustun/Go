package main

import "fmt"

/*
Collection of disparate data types that describe a single concept
anonymous structs are allowed (aDoctor := struct{name string}{name: "Ozgur Ustun"})
Structs are value types. When you assign one struct variable to another, a new copy of the struct is created and assigned.
Similarly, when you pass a struct to another function, the function gets its own copy of the struct.
*/

type Doctor struct {
	number     int
	actorName  string
	companions []string
	episodes   []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Smith",
		},
	}

	/*
		Also can be created without field names, but not adviced

		aDoctor := Doctor{
				3,
				"Jon Pertwee",
				[]string{
					"Liz Shaw",
					"Jo Grant",
					"Sarah Smith",
				},
			}
	*/

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])
	fmt.Println(aDoctor.episodes)
}
