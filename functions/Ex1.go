package main

import "fmt"

/*

	Basic syntax:
		func asd(){
		...
		}

	Parameters:
		- Comma delimited list of variables and types

			func asd(arg1 string,arg2 string) {
			}

		- Parameters of same type list type once

			func asd(arg1,arg2 string){
			}

		- If pointers are passed in, function can change the value in the caller
		  	For slices and maps this will be happened even if we not pass the pointer of slices or maps to the function. They work with internal pointers

		- Use variadic parameters to send list of same types in
			Must be last parameter
			Received as slice
			func foo(bar string, baz ...int){
			}

	Return values:
		- Single return values just list type
			func foo() int{
			}
		- Multiple return value list types surrounded by parantheses
			func foo() (int,error)
			The (result type, error) paradigm is very common idiom
		- Can use named return values
			Initializes returned variable
			Return using return keywork on its own
			func result(args ...int) (result int) {
				for _, v := range args {
				result += v
				}
				return
			}
		- Can return address of local variables
			Automatically promoted from local memory(stack) to shared memory(heap)

	Anonymous functions
		- Functions don't have names
		- Immediately invoked
			func() {
				...
			}()
		- Assigned to a variabl or passed as an argument to a function
			a := func () {
				...
			}
			a()

	Function as types
		- Can assign functions to variables or use as arguments and return values in functions
		- Type signature is like function signature, with no names
			var functionAssignedVariable func(string,string,int) (int,error)

	Methods
		- Function that executes in context of a type
		- Format
			func (g greeter) greet(){
				...
			}
		- Receiver can be value or pointer
			Value receiver gets copy of type
			Pointer receiver gets pointer of type
*/

func main() {
	arg1 := "asd"
	arg2 := "def"
	func1(arg1, arg2)
	fmt.Println(arg2)
}

func func1(arg1, arg2 string) {
	fmt.Println(arg1, arg2)
	arg2 = "ghj"
	fmt.Println(arg2)
}
