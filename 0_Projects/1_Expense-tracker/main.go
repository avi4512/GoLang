package main

import "fmt"

func main() {

	loadFromJson()
	for {
		menu()
		var choice int
		fmt.Print("Enter your Choice (1-5):")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			addExpenses()
		case 2:
			viewExpenses()
		case 3:
			totalExpenses()
		case 4:
			categoryWiseExpenses()
		case 5:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid choice...!\n choose Again:")
			continue
		}

	}

}

func menu() {
	fmt.Println("===== Student Expense Tracker =====")
	fmt.Println("1. Add new expense")
	fmt.Println("2. View all expenses")
	fmt.Println("3. View total expense")
	fmt.Println("4. View category-wise expense")
	fmt.Println("5. Exit")
}
