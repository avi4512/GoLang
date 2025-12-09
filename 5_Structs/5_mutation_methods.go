package main

import (
	"fmt"
	"time"
)

//Defining Struct

type userData struct {
	firstName string
	lastName  string
	birthDate string
	createdOn time.Time
}

func main() {

	userFirstName := getData("Enter your First Name: ")
	userLastName := getData("Enter your Last Name: ")
	userBirthDate := getData("Enter your Data of Birth DD/MM/YYYY: ")

	appUser := userData{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdOn: time.Now(),
	}

	appUser.outputDetails()
	appUser.clearNames()
	appUser.outputDetails()

}

func (u *userData) outputDetails() {

	println(u.firstName, u.lastName, u.birthDate)
	//println((*u).firstName, (*u).lastName, (*u).birthDate) other way
}

func (u *userData) clearNames() {

	u.firstName = ""
	u.lastName = ""
}

func getData(msgText string) string {
	fmt.Print(msgText)
	var data string
	fmt.Scan(&data)
	return data
}
