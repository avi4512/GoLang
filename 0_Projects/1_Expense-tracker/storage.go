package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var storage_expenses = make([]expenses, 0)

func saveToJson(se *[]expenses) {

	byte_data, err := json.Marshal(*se)

	if err != nil {
		fmt.Println("File not saved")
		fmt.Println(err)
	}

	os.WriteFile("Expenses.json", byte_data, 0644)

	path, _ := os.Getwd()
	fmt.Printf("File Created Successfully..\nat %s\n", path)
}

func loadFromJson() {

	fileName := "expenses.json"

	byte_data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Printf("%s: such file Not Found!!\n", fileName)
		return
	}

	json.Unmarshal(byte_data, &storage_expenses)

}
