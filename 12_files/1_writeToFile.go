package main

import (
	"fmt"
	"os"
)

func writeToFile(text string) {
	fmt.Print(text)
	var amount float64
	fmt.Scan(&amount)

	str_amount := fmt.Sprint(amount)
	os.WriteFile("amount_file.txt", []byte(str_amount), 0644)
	path, _ := os.Getwd()
	fmt.Println("Success.....")
	fmt.Println("File saved at", path)
}

func main() {
	writeToFile("Enter the Amount:")
}
