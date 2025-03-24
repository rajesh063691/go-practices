package main

import "fmt"

const (
	a = 1 << (iota + 10)
	b
	//tetststs
	c
	_ // skip value
	d
)

// const (
// 	a = iota + 10
// 	b = iota + 2
// 	c = iota + 3
// 	d = iota + 5
// )

// Authorization Flags
const (
	// Caller Types
	FrameworkAuthInternalAllowed = 1 << (iota * 10)
	FrameworkAuthRootAllowed
	FrameworkAuthServiceProviderAllowed
	FrameworkAuthTenantAllowed

	// Calling Requirements
	FrameworkAuthTenantRequired
	FrameworkAuthOrgRequired

	// Permissions Required
	FrameworkAuthOrgOwnerPermission
	FrameworkAuthTenantAdminPermission

	// Credential Types Accepted
	FrameworkAuthPasswordAccepted
	FrameworkAuthBearerTokenAccepted
	FrameworkAuthServiceTokenAccepted
	FrameworkAuthInitKeyAccepted
)

func main() {
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)

	fmt.Printf("%v \n", FrameworkAuthInternalAllowed)
	fmt.Printf("%v \n", FrameworkAuthRootAllowed)
	fmt.Printf("%v \n", FrameworkAuthServiceProviderAllowed)
}
