package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type expenses struct {
	Title    string  `json:"title"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
}

func main() {

	expSlice := make([]expenses, 0)

	exp1 := newExpenses("soap", 45.7, "house")

	exp2 := expenses{
		Title:    "Lunch",
		Amount:   120,
		Category: "Food",
	}
	exp3 := expenses{
		"bus ticket",
		145,
		"tickets",
	}

	expSlice = append(expSlice, exp1)
	expSlice = append(expSlice, exp2)
	expSlice = append(expSlice, exp3)

	for index, val := range expSlice {
		fmt.Printf("%d. Title: %s | Amount: %.2f | Category: %s\n", index+1, val.Title, val.Amount, val.Category)
		fmt.Println("------------------------------")
	}
	fmt.Println(expSlice)

	saveToJson(expSlice)

}

func newExpenses(title string, amount float64, category string) expenses {

	return expenses{
		Title:    title,
		Amount:   amount,
		Category: category,
	}
}

func saveToJson(sli []expenses) {

	byte_data, _ := json.Marshal(sli)

	os.WriteFile("expense-tracker.json", byte_data, 0644)
	fmt.Println("file created successfully")
}
