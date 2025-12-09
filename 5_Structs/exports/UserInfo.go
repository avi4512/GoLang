package user

import (
	"errors"
	"time"
)

//Defining Struct

type UserData struct {
	firstName string
	lastName  string
	birthDate string
	createdOn time.Time
}

// Constructor With Validation's
func New(FirstName, LastName, BirthDate string) (*UserData, error) {

	if FirstName == "" || LastName == "" || BirthDate == "" {
		return nil, errors.New("required FirstName LastName and BirthDate")

	}
	return &UserData{
		firstName: FirstName,
		lastName:  LastName,
		birthDate: BirthDate,
		createdOn: time.Now(),
	}, nil
}

func (u *UserData) OutputDetails() {

	println(u.firstName, u.lastName, u.birthDate)
	//println((*u).firstName, (*u).lastName, (*u).birthDate) other way
}

func (u *UserData) ClearNames() {

	u.firstName = ""
	u.lastName = ""
}
