package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// embeded struct
type Admin struct {
	email    string
	password string
	// you can pass a name like
	// user User
	// but you can do it anonimusly like this and keep in mind if you pass a custom name that all with lowercase you can not access to other method and items in that User struct and you shuld always use this name like admin.user.firstname
	User
}

func NewAdmin(email, passwrod string) Admin {
	return Admin{
		email:    email,
		password: passwrod,
		User: User{
			firstName: "admin",
			lastName:  "admin",
			birthdate: "----",
			createdAt: time.Now(),
		},
	}
}

// you can avoid copy
func New(firstName, lastName, birthdate string) (*User, error) {
	// validate create user in constuctor
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstName and lastName and birthdate are reaquired")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func (u User) OutputUserDetail() {
	fmt.Printf(` username is %s and the userlastName is %s,this user born in %s. this info created at %s`, u.firstName, u.lastName, u.birthdate, u.createdAt.Local())
}
func (u *User) ClearUser() {
	u.firstName = ""
	u.lastName = ""
}
