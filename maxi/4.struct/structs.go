package main

import (
	"fmt"

	"example.com/struct/user"
)

// type user struct {
// 	firstName string
// 	lastName  string
// 	birthdate string
// 	createdAt time.Time
// }

// // you can avoid copy
// func newUser(firstName, lastName, birthdate string) (*user, error) {
// 	// validate create user in constuctor
// 	if firstName == "" || lastName == "" || birthdate == "" {
// 		return nil, errors.New("firstName and lastName and birthdate are reaquired")
// 	}
// 	return &user{
// 		firstName: firstName,
// 		lastName:  lastName,
// 		birthdate: birthdate,
// 		createdAt: time.Now(),
// 	}, nil
// }

// func (u user) outputUserDetail() {
// 	fmt.Printf(` username is %s and the userlastName is %s,this user born in %s. this info created at %s`, u.firstName, u.lastName, u.birthdate, u.createdAt.Local())
// }
// func (u *user) clearUser() {
// 	u.firstName = ""
// 	u.lastName = ""
// }

type str string

func (text str) log() {
	fmt.Println(text)
}
func main() {
	var test str = "test"
	test.log()
	fmt.Println("from fmt: ", test)
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!
	// var userData user
	// userData = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthdate: userBirthdate,
	// 	createdAt: time.Now(),
	// }
	var userData *user.User
	userData, err := user.New(userFirstName, userLastName, userBirthdate)

	adminData := user.NewAdmin("fardin@gmail.com", "1234")
	adminData.OutputUserDetail()

	if err != nil {
		fmt.Println(err)
		return
	}

	// outputUserDetail(&userData)
	userData.OutputUserDetail()
	userData.ClearUser()
	userData.OutputUserDetail()

}

// func outputUserDetail(user *user) {
// 	fmt.Printf(`username is %s and the userlastName is %s,this user born in %s. this info created at %s`, user.firstName, user.lastName, user.birtyear, user.createdAt.Local())
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	// for accept newline Scanln
	fmt.Scanln(&value)
	return value
}
