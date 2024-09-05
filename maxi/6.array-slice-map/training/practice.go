package main

import (
	"fmt"
)

type Product struct {
	title string
	id    int
	price float64
}

func New(title string, id int, price float64) *Product {
	return &Product{
		title: title,
		id:    id,
		price: price,
	}
}

func main() {
	hobbise := [3]string{"gameing", "learning", "music"}

	fmt.Println(hobbise)

	firsthobby := hobbise[0]
	fmt.Println(firsthobby)

	// resthobby := hobbise[1:]
	// resthobby := hobbise[1:3]
	resthobby := [2]string{hobbise[1], hobbise[2]}
	fmt.Println(resthobby)

	// first2Slice := hobbise[0:2]
	first2Slice := hobbise[:2]
	fmt.Println(first2Slice)

	// first2Slice[0] = hobbise[1]
	// first2Slice[1] = hobbise[2]
	// fmt.Println(first2Slice)
	// fmt.Println(hobbise)
	// i dont understanding question so the right thing is this its about cap() :
	first2Slice = first2Slice[1:3]
	fmt.Println(first2Slice)
	fmt.Println(hobbise)

	goals := []string{"learn_go", "build reat api without fraimwork"}
	fmt.Println(goals)
	goals[1] = "be a better go developer"
	goals = append(goals, "be a better software engineer")
	fmt.Println(goals)

	fmt.Println("_________________")

	// firstProduct := New("asus", 1, 199.99)
	// secondProduct := New("a24", 2, 99.99)

	// products := []Product{*firstProduct, *secondProduct}

	// therthProduct := New("x3", 3, 88.99)
	// products = append(products, *therthProduct)
	// fmt.Println(products)

	// or
	// products := []Product{Product{"asus", 1, 199.99}, Product{"a24", 2, 99.99}}
	// you can skip puting Product for createa a Product go know that in this array is a struct
	products := []Product{{"asus", 1, 199.99}, {"a24", 2, 99.99}}
	products = append(products, Product{"x3", 3, 88.99})
	fmt.Println(products)

}
