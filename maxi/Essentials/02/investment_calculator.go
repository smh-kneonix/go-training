package main

//section 1

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var investmentAmount = 1000
// 	var expectReturnRate = 5.5
// 	var year = 10
// 	var futureValue = float64(investmentAmount) * math.Pow(1+expectReturnRate/100, float64(year))

// 	fmt.Println(futureValue)

// }

//section 2

// import (
// 	"fmt"
// 	"math"
// )

// // for global scoop to use in main and other function in this file
// const inflationValue = 2.5

// func main() {
// 	var investmentAmount float64 = 1000
// 	// with printf you can create a formatted string and put your value in every where you want in your string by using %g or %T(%g put exact value of your varible and %T put type or your varible)
// 	fmt.Printf("type: %T and value: %g \n", investmentAmount, investmentAmount)
// 	var expectReturnRate = 5.5
// 	year := 10.0

// 	fmt.Printf("the pointer addresss of year is %p \n", &year)

// 	// fmt.Print("Investment Amount: ")
// 	outputText("Investment Amount: ")
// 	fmt.Scan(&investmentAmount)

// 	// fmt.Print("Expect Return Rate: ")
// 	outputText("Expect Return Rate: ")
// 	fmt.Scan(&expectReturnRate)

// 	// fmt.Print("Year: ")
// 	outputText("Year: ")
// 	fmt.Scan(&year)

// 	// replace with function calculate
// 	// futureValue := investmentAmount * math.Pow(1+expectReturnRate/100, year)
// 	// futureRealValue := futureValue / math.Pow(1+inflationValue/100, year)

// 	futureValue, futureRealValue := calculatoFutureValues(investmentAmount, expectReturnRate, year)
// 	fmt.Println(futureValue)
// 	fmt.Println(futureRealValue)

// 	var input = 5.5
// 	fmt.Scan(&input)
// 	fmt.Println(input, "nice")

// 	var num1, num2 string = "1000", "10"
// 	fmt.Println(num1, num2)

// }
// func calculatoFutureValues(investmentAmount, expectReturnRate, year float64) (float64, float64) {
// 	fv := investmentAmount * math.Pow(1+expectReturnRate/100, year)
// 	rfv := fv / math.Pow(1+inflationValue/100, year)
// 	return fv, rfv
// }

// // you can do it like this too, but just do it for small function because it's hard to read
// // func calculatoFutureValues(investmentAmount, expectReturnRate, year float64) (fv float64, rfv float64) {
// // 	fv = investmentAmount * math.Pow(1+expectReturnRate/100, year)
// // 	rfv = fv / math.Pow(1+inflationValue/100, year)
// // 	// return fv, rfv
// // 	// or
// // 	return
// // }

// func outputText(text string) {
// 	print(text)
// }

// traning

func main() {
	// see profit-calculator.go
	ProfitCalculator()
}

// section 3

// func main() {

// }
