package main

import "fmt"

func main() {

	// creation of maps

	product := map[int]string{
		101: "Soap",
		102: "Paste",
		103: "Shampoo",
	}
	fmt.Println(product)
	fmt.Println("-------------")

	// Accessing Values using key
	fmt.Println("Accessing value using key:", product[102])

	// adding or changing value based on key
	product[104] = "Conditionar"
	fmt.Println(product)

	// delete key
	delete(product, 102)
	fmt.Println("After deletion:", product)
}
