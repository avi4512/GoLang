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

func writeToExistingFile(text string) {

	fmt.Print(text)
	var amount float64
	fmt.Scan(&amount)

	file, err := os.OpenFile("amount_file.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	str_amount := fmt.Sprintf("%v\n", amount)
	file.WriteString(str_amount)

	path, _ := os.Getwd()
	fmt.Println("added Success.....")
	fmt.Println("File saved at", path)

}

func main() {
	writeToExistingFile("Enter the Amount:")
}
