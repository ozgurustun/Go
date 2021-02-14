package main

import "fmt"

func main() {
	/*
		Defer:
		After function finish its final statement, but before it actually returns
		LIFO order
		Used to delay execution of a statement until function exists
		Useful to group open and close functions together
		Arguments evaluated at time defer is executed, not at time of called function execution

		Panic:
		Occur a program can not continue at all
		Don't use a file can not be opened, unless it is critical
		Use for unrecovable events - can not obtain TCP port for Web Server
		function will stop executing but deferred functions will still fire
		if nothing handles panic, program will exit

		Recover:
		Only useful in deferred functions
		Used to recover from panics
		Current function will not attempt to continue, but higher functions in call stack will

	*/

	fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
}
