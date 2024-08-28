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

import (
	"fmt"
	"math"
)

func main() {
	const inflationValue = 2.5
	var investmentAmount float64 = 1000
	// with printf you can create a formatted string and put your value in every where you want in your string by using %g or %T(%g put exact value of your varible and %T put type or your varible)
	fmt.Printf("type: %T and value: %g \n", investmentAmount, investmentAmount)
	var expectReturnRate = 5.5
	year := 10.0
	futureValue := investmentAmount * math.Pow(1+expectReturnRate/100, year)

	futureRealValue := futureValue / math.Pow(1+inflationValue/100, year)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

	var input = 5.5
	fmt.Scan(&input)
	fmt.Println(input, "nice")

	var num1, num2 string = "1000", "10"
	fmt.Println(num1, num2)

}
