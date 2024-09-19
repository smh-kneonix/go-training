package main

import (
	"fmt"
	"math"
)

func main() {
	var i int = 42
	fmt.Println(float32(i))
	fmt.Println(float64(i))
	fmt.Println(uint(i))
	fmt.Println(string(i))

	fmt.Println(math.Sqrt(float64(i))) // you should change int to float64

	// const
	const hi = "hi"
	// hi = "bye" //error

	// numeric constant
	const big = 1 << 100 // make 1 with 100 zero after that in binary
	const small = big >> 99 // delete 99 number from right in binary
	// fmt.Println(small) // you can not log this because its too much big
	fmt.Println(small) // 2

	// for
	var sum int
	for i:=0; i < 10; i++ {
		sum += i
	}

	// if
	num := 2	
	if num==1 {
		fmt.Println(1)
	} else if num > 2 {
		fmt.Println(2)
	} else {
		fmt.Println(0)
	}
	
}