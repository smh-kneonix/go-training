package main

import "fmt"

func main() {
	prices := []float64{36.99, 41.99, 1.99}

	fmt.Println(prices[1])
	prices[2] = 25.99 // ok!
	// prices[3] = 44.99 // panic: runtime error: index out of range [3] with length 3

	prices = append(prices, 78.99, 55.99, 21.99)
	fmt.Println(prices)

	discountPrices := []float64{27.99, 17.99, 38.99}

	prices = append(prices, discountPrices...)
	fmt.Println(prices)

	websites := map[string]string{
		"google":             "https://google.com",
		"amazon web service": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["amazon web service"])

	websites["linkedIn"] = "https://linkedIn.com"
	fmt.Println(websites)

	delete(websites, "google")
	fmt.Println(websites)
}

// func main() {
// 	// prices := [5]float64{99, 36.99, 41.99, 24.99, 12.99}
// 	// fmt.Println(prices)

// 	// featurePrice := prices[1:5]
// 	// fmt.Println(featurePrice)
// 	prices := [5]float64{99, 36.99, 41.99, 24.99, 12.99}
// 	fmt.Println(prices)

// 	featurePrice := prices[1:]
// 	featurePrice[0] = 192.99

// 	fmt.Println(prices)

// 	arr := [5]int{1, 2, 3, 4, 5}
// 	slice1 := arr[1:]
// 	slice2 := slice1[:1]
// 	fmt.Println(len(slice1), cap(slice1))
// 	fmt.Println(len(slice2), cap(slice2))

// 	slice2 = slice2[:4]
// 	fmt.Println(slice2)
// 	fmt.Println(len(slice2), cap(slice2))
// }
