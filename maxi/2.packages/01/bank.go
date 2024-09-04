package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Print("Wellcome to go bank!")
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
		fmt.Println("_____________")
		// panic("Can't continue, sorry.")
	}
	for {
		PercentOption()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance updated! New amount: ", accountBalance)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. Must be less than you account balance!")
				continue
			}
			accountBalance -= withdrawalAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance updated! New amount: ", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return

		}
	}
}
