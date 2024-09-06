package main

import "fmt"

func main() {
	age := 32

	agePointer := &age

	fmt.Println("age: ", *agePointer)

	adultYear := getAdultYear(agePointer)

	fmt.Println("adalutYear: ", adultYear)

	editAgeToAdultYear(agePointer)
	fmt.Println("after edit: ", age)
}

func getAdultYear(age *int) int {
	return *age - 18
}
func editAgeToAdultYear(age *int) {
	*age = *age - 18
}
