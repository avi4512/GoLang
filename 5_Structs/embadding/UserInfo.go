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

type Admin struct {
	email    string
	password string
	UserData
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

// Admin Struct Constructor
func NewAdmin(email, password string) Admin {

	return Admin{
		email:    email,
		password: password,
		UserData: UserData{
			firstName: "ABC",
			lastName:  "XYZ",
			birthDate: "01/01/2000",
			createdOn: time.Now(),
		},
	}

}

func (u *UserData) OutputDetails() {

	println(u.firstName, u.lastName, u.birthDate)
	//println((*u).firstName, (*u).lastName, (*u).birthDate) other way
}

func (u *UserData) ClearNames() {

	u.firstName = ""
	u.lastName = ""
}
