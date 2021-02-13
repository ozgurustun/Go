package main

//Capital letter exported from the package
var I int = 33

func main() {

	//Naming conventions controls visibility
	//lowercase variables are scoped to the package,
	//which means anything consume the package cant see and manipulate variables with lowercase

	var i int = 3
	var theURL = "https://google.com"
	var theHTTPRequest = "https://google.com"
}
