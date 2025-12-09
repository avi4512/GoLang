package main

import "fmt"

func main() {
	age := 23

	agePointer := &age

	fmt.Println("Age:", age)
	fmt.Println("Age Pointer:", agePointer)

	adultYears := getAdultAge(agePointer)
	fmt.Println("Adult Yeras:", adultYears)
}

func getAdultAge(agePointer *int) int {
	adultAge := *agePointer - 18
	return adultAge

}
