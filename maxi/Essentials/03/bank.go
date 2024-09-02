package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// my code
// func main() {
// 	fmt.Println("welcome to bank")
// 	var action int
// 	var amount float64
// 	balance := 0.0

// 	for action != 4 {
// 		fmt.Println("1. check balance")
// 		fmt.Println("2. deposit")
// 		fmt.Println("3. witdrow")
// 		fmt.Println("4. exit")
// 		fmt.Println("please choose a number of command!")
// 		fmt.Print("action: ")
// 		fmt.Scan(&action)
// 		if action == 1 {

// 			fmt.Printf("your balance is: %.2f \n", balance)
// 		} else if action == 2 {
// 			fmt.Print("amount of deposit: ")
// 			fmt.Scan(&amount)
// 			if amount < 0 {
// 				fmt.Println("invalid amount, the amount must be greater than 0")
// 				return
// 			}
// 			balance += amount
// 		} else if action == 3 {
// 			fmt.Print("amount of witdrow: ")
// 			fmt.Scan(&amount)
// 			if amount < 0 {
// 				fmt.Println("invalid amount, the amount must be greater than 0")
// 				return
// 			}
// 			if amount > balance {
// 				fmt.Println("you can not witdrow more than your balance")
// 				return
// 			}
// 			balance -= amount
// 		} else if action == 4 {
// 			fmt.Println("see you soon!")
// 		}

// 	}
// }

// maxi code
// const accountBalanceFile = "balance.txt"

// func getBalanceFromFile() (float64, error) {
// 	data, err := os.ReadFile(accountBalanceFile)
// 	if err != nil {
// 		return 1000, errors.New("can not read file, so you can go with 1000 bounus!")
// 	}
// 	balanceText := string(data)
// 	balance, err := strconv.ParseFloat(balanceText, 64)
// 	if err != nil {
// 		return 1000, errors.New("can not convert string to float, so you can go with 1000 bounus!")
// 	}
// 	return balance, nil
// }

// func writeBalanceToFile(balance float64) {
// 	var balanceText = fmt.Sprint(balance)
// 	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
// }

// func main() {
// 	fmt.Print("Wellcome to go bank!")
// 	var accountBalance, err = getBalanceFromFile()
// 	if err != nil {
// 		fmt.Println("ERROR:")
// 		fmt.Println(err)
// 		fmt.Println("_____________")
// 		// panic("Can't continue, sorry.")
// 	}
// 	for {
// 		fmt.Println("What do you want to do?")
// 		fmt.Println("1. Check balance")
// 		fmt.Println("2. Deposit money")
// 		fmt.Println("3. Withdraw modey")
// 		fmt.Println("4. Exit")
// 		var choice int
// 		fmt.Print("Your choice: ")
// 		fmt.Scan(&choice)

// 		switch choice {
// 		case 1:
// 			fmt.Println("Your balance is: ", accountBalance)
// 		case 2:
// 			fmt.Print("Your deposit: ")
// 			var depositAmount float64
// 			fmt.Scan(&depositAmount)
// 			if depositAmount <= 0 {
// 				fmt.Println("Invalid amount. Must be greater than 0.")
// 				continue
// 			}
// 			accountBalance += depositAmount
// 			writeBalanceToFile(accountBalance)
// 			fmt.Println("Balance updated! New amount: ", accountBalance)
// 		case 3:
// 			fmt.Print("Withdrawal amount: ")
// 			var withdrawalAmount float64
// 			fmt.Scan(&withdrawalAmount)
// 			if withdrawalAmount > accountBalance {
// 				fmt.Println("Invalid amount. Must be less than you account balance!")
// 				continue
// 			}
// 			accountBalance -= withdrawalAmount
// 			writeBalanceToFile(accountBalance)
// 			fmt.Println("Balance updated! New amount: ", accountBalance)
// 		default:
// 			fmt.Println("Goodbye!")
// 			fmt.Println("Thanks for choosing our bank")
// 			return

// 		}
// 	}
// }
