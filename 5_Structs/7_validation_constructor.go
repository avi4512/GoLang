package main

import (
	"errors"
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

// Constructor With Validation's
func newUserData(FirstName, LastName, BirthDate string) (*userData, error) {

	if FirstName == "" || LastName == "" || BirthDate == "" {
		return nil, errors.New("required FirstName LastName and BirthDate")

	}
	return &userData{
		firstName: FirstName,
		lastName:  LastName,
		birthDate: BirthDate,
		createdOn: time.Now(),
	}, nil
}

func main() {

	userFirstName := getData("Enter your First Name: ")
	userLastName := getData("Enter your Last Name: ")
	userBirthDate := getData("Enter your Data of Birth DD/MM/YYYY: ")

	appUser, err := newUserData(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		panic("Somthing Went Wrong CheckOut Once")
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
	fmt.Scanln(&data)
	return data
}
