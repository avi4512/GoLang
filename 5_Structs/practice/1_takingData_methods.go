package main

import "fmt"

type user struct {
	firstName string
	lastName  string
}

func (u user) outputDetails() {
	fmt.Println(u.firstName, u.lastName)
}

func (u *user) getDetails(fn, ln string) {
	u.firstName = fn
	u.lastName = ln

}

func main() {

	var appuser user
	var fn, ln string
	fmt.Print("Enter first name:")
	fmt.Scan(&fn)
	fmt.Print("Enter last name:")
	fmt.Scan(&ln)
	appuser.getDetails(fn, ln)
	appuser.outputDetails()

}
