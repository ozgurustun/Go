package main

/*

Using tag

*/

import (
	"fmt"
	"reflect"
)

type Teacher struct {
	Name string `required max:"100"`
	Age  int
}

func main() {

	//Get struct`s field tag
	a := reflect.TypeOf(Teacher{})
	field, _ := a.FieldByName("Name")
	fmt.Println(field.Tag)

}
