package tasks

import (
	"errors"
	"fmt"
	"os"
)

func ProfitCalculator() {

	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		panic(err)
	}
	expenses, err := getUserInput("Rxpenses: ")
	if err != nil {
		panic(err)
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		panic(err)
	}

	EBT, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.0f, profit: %.0f, ratio: %.2f", EBT, profit, ratio)
}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("invalid input, input must be greater than 0")
	}
	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	resultText := fmt.Sprint(ebt, profit, ratio)

	os.WriteFile("calculateFinancials.txt", []byte(resultText), 0644)
	return ebt, profit, ratio
}
