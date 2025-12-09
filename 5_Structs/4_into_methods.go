package main

import "fmt"

type user struct {
	firstName string
	lastName  string
}

func (u user) outputDetails() {
	fmt.Println(u.firstName, u.lastName)
}

func main() {

	appUser := user{"Avi", "Reddy"}
	appUser.outputDetails()

}
