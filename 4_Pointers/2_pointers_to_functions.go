package main

import "fmt"

func main() {
	age := 23

	agePointer := &age

	fmt.Println("Age:", age)
	fmt.Println("Age Pointer:", agePointer)

	getAdultAge(agePointer)
}

func getAdultAge(agePointer *int) {
	adultAge := *agePointer - 18
	fmt.Println("Adult Age years:", adultAge)
}
