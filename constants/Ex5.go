package main

import "fmt"

const (
	_  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func main() {
	fileSize := 4000000000.
	fmt.Printf("%.2fGB", fileSize/GB)
}
