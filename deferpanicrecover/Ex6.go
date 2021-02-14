package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("starting panic...")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}

func panickerUnhandled() {
	fmt.Println("starting panic...")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			panic(err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}
