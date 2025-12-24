package main

import (
	"errors"
	"fmt"
	"strconv"
)

type expenses struct {
	Title    string  `json:"title"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
}

func NewExpenses(title string, amount float64, category string) (expenses, error) {
	if title == "" || category == "" || amount <= 0 {
		return expenses{}, errors.New("error: Must Fill all Details")
	}
	return expenses{
		Title:    title,
		Amount:   amount,
		Category: category,
	}, nil
}

func getData(msgText string) string {
	fmt.Print(msgText)
	var data string
	fmt.Scan(&data)
	return data
}

func addExpenses() {

	title := getData("Enter the title:")
	amount := getData("Enter the amount:")
	amount_float, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		fmt.Println("error:Enter valid number")
		fmt.Println(err)
		return
	}
	category := getData("Enter the Category:")
	exp, _ := NewExpenses(title, amount_float, category)
	storage_expenses = append(storage_expenses, exp)
	saveToJson(&storage_expenses)
	fmt.Println("Expense added successfully!")

}

func viewExpenses() {
	fmt.Println("----------All Expenses-----------")

	if len(storage_expenses) == 0 {
		fmt.Println("No expenses found")
		return
	}

	for i, exp := range storage_expenses {
		fmt.Printf("%d. %s | %.2f | %s\n",
			i+1, exp.Title, exp.Amount, exp.Category)
	}

}

func totalExpenses() {

	if len(storage_expenses) == 0 {
		fmt.Println("No expenses available to calculate total.")
	}
	total := 0.0
	for _, val := range storage_expenses {
		total = total + val.Amount
	}
	fmt.Printf("Total Expense: ₹%.2f\n", total)
}

func categoryWiseExpenses() {

	var expenseMap = make(map[string]float64, 0)
	for _, val := range storage_expenses {

		expenseMap[val.Category] += val.Amount
	}
	for keys, val := range expenseMap {
		fmt.Println(keys, ":", val)
	}

}
