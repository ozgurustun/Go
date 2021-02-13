package main

import "fmt"

const (
	isAdmin          = 1 << iota // 0000 0001
	isHeadquarters               // 0000 0010
	canSeeFinancials             // 0000 0100

	canSeeAfrica       // 0000 1000
	canSeeAsia         // 0001 0000
	canSeeEurope       // 0010 0000
	canSeeNorthAmerica // 0100 0000
	canSeeSouthAmerica // 1000 0000
)

func main() {
	//roles 100101
	var roles = isAdmin | canSeeFinancials | canSeeEurope
	//bitmasking
	fmt.Printf("%b\n", roles)
	// isAdmin & roles return 000001
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	//isHeadquarters & roles return 000010
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)
}
