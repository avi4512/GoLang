package main

import (
	"GoLang/fileOprs"
	"fmt"
)

const globalAmountFile = "balance.txt"

func main() {

	currentBalance, err := fileOprs.ReadFloatFromFile(globalAmountFile)

	if err != nil {
		fmt.Println("Something went Worng..!")
		fmt.Println(err)
		//panic("Program got Terminated..!!!!")
	}
	fmt.Println("Welcome the Go Banking..")

	for {

		chooseOption()

		var choice int
		fmt.Print("Enter your Choice: ")
		fmt.Scan(&choice)
		switch choice {

		case 1:
			fmt.Printf("Current Bank Balance : %.2f\n", currentBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter amount to Deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount > 0 {
				currentBalance += depositAmount
				fmt.Printf("After Deposit Amount Added: %.2f\n", currentBalance)
				fileOprs.WriteFloatToFile(globalAmountFile, currentBalance)
			} else {
				fmt.Println("Deposit Amount must be GREATER than 0")
				continue
			}
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter amount to Deposit: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > 0 && withdrawAmount <= currentBalance {
				currentBalance -= withdrawAmount
				fmt.Printf("After withdraw: %.2f\n", currentBalance)
				fileOprs.WriteFloatToFile(globalAmountFile, currentBalance)
			} else {
				fmt.Println("INVALID withdraw Amount...!")
				continue
			}
		default:
			fmt.Println("Thank You..")
			fmt.Println("For Choosing Go Bank...")
			return
		}
	}
}
