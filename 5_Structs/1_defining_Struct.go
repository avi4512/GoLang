package main

import "fmt"

//Defining Struct

type userData struct {
	fistName  string
	lastName  string
	birthDate string
}

func main() {

	firstName := getData("Enter your First Name: ")
	lastName := getData("Enter your Last Name: ")
	birthDate := getData("Enter your Data of Birth DD/MM/YYYY: ")

	// fmt.Println(firstName)
	// fmt.Println(lastName)
	// fmt.Println(birthDate)

	output(firstName, lastName, birthDate)

}

func output(fist_Name, last_Name, date_of_birth string) {

	print(fist_Name, last_Name, date_of_birth)
}

func getData(msgText string) string {
	fmt.Print(msgText)
	var data string
	fmt.Scan(&data)
	return data
}
