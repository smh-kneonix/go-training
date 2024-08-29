package main

import "fmt"

func main() {
	fmt.Println("welcome to bank")
	var action int
	var amount float64
	balance := 0.0

	for action != 4 {
		fmt.Println("1. check balance")
		fmt.Println("2. deposit")
		fmt.Println("3. witdrow")
		fmt.Println("4. exit")
		fmt.Println("please choose a number of command!")
		fmt.Print("action: ")
		fmt.Scan(&action)
		if action == 1 {

			fmt.Printf("your balance is: %.2f \n", balance)
		} else if action == 2 {
			fmt.Print("amount of deposit: ")
			fmt.Scan(&amount)
			if amount < 0 {
				fmt.Println("invalid amount, the amount must be greater than 0")
				return
			}
			balance += amount
		} else if action == 3 {
			fmt.Print("amount of witdrow: ")
			fmt.Scan(&amount)
			if amount < 0 {
				fmt.Println("invalid amount, the amount must be greater than 0")
				return
			}
			if amount > balance {
				fmt.Println("you can not witdrow more than your balance")
				return
			}
			balance -= amount
		} else if action == 4 {
			fmt.Println("see you soon!")
		}

	}
}
