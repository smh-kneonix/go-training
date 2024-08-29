package main

import "fmt"

// func ProfitCalculator() {
// 	var revenue, expenses, taxRate float64

// 	// fmt.Print("Revenue: ")
// 	// fmt.Scan(&revenue)
// 	getUserInput("Revenue", &revenue)

// 	// fmt.Print("Rxpenses: ")
// 	// fmt.Scan(&expenses)
// 	getUserInput("Rxpenses", &expenses)

// 	// fmt.Print("Tax Rate: ")
// 	// fmt.Scan(&taxRate)
// 	getUserInput("Tax Rate", &taxRate)

// 	// EBT := revenue - expenses
// 	// profit := EBT * (1 - taxRate/100)
// 	// ratio := EBT / profit
// 	var EBT, profit, ratio float64 = ebtProfitRatioCalcultor(revenue, expenses, taxRate)

// 	// %f is a place holder for desimal number and you can controll that numbers you can reed more https://pkg.go.dev/fmt#hdr-Printing
// 	fmt.Printf("EBT: %.0f, profit: %.0f, ratio: %.2f", EBT, profit, ratio)
// }

// func getUserInput(text string, x *float64) {
// 	fmt.Printf(`%s: `, text)
// 	fmt.Scan(x)
// }

// func ebtProfitRatioCalcultor(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
// 	ebt = revenue - expenses
// 	profit = ebt * (1 - taxRate/100)
// 	ratio = ebt / profit
// 	return ebt, profit, ratio
// }

// the code below was my code but the maximilan code is here i commented this but you can comment the code below and uncomment maxi code
// maxi code:

func ProfitCalculator() {
	// var revenue, expenses, taxRate float64
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Rxpenses: ")
	taxRate := getUserInput("Tax Rate: ")

	EBT, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.0f, profit: %.0f, ratio: %.2f", EBT, profit, ratio)
}

func getUserInput(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
