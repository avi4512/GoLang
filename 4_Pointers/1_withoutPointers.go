package main

import "fmt"

func main() {

	age := 23

	//var agePointer *int
	agePointer := &age

	fmt.Println("Age:", age)
	fmt.Println("Age Pointer:", agePointer)
	getAdultAge(age)
}

func getAdultAge(age int) {

	adultAge := age - 18
	fmt.Println("Adult Age yeras:", adultAge)

}
