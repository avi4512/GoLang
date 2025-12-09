package main

import "fmt"

func main() {
	age := 23

	agePointer := &age

	fmt.Println("Age:", age)
	fmt.Println("Age Pointer:", agePointer)

	editAgeToAdultAge(agePointer)
	fmt.Println("Age After manulpulation:", age)
}

func editAgeToAdultAge(agePointer *int) {
	*agePointer = *agePointer - 18

}
