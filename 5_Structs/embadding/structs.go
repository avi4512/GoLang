package main

import (
	"GoLang/user"
	"fmt"
)

func main() {

	userFirstName := getData("Enter your First Name: ")
	userLastName := getData("Enter your Last Name: ")
	userBirthDate := getData("Enter your Data of Birth DD/MM/YYYY: ")

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	adminUser := user.NewAdmin("avi.punyala@gmail.com", "xxxxxxxxxxx")

	adminUser.OutputDetails()
	adminUser.ClearNames()
	adminUser.OutputDetails()

	if err != nil {
		fmt.Println(err)
		panic("Somthing Went Wrong CheckOut Once")
	}

	appUser.OutputDetails()
	appUser.ClearNames()
	appUser.OutputDetails()

}

func getData(msgText string) string {
	fmt.Print(msgText)
	var data string
	fmt.Scanln(&data)
	return data
}
