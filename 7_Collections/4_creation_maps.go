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
	fmt.Println(product[102])
}
