package user

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

// struct embedding

type Admin struct {
	email    string
	password string
	User
}

// structs methods

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthday)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthday:  "---",
			createdAt: time.Now(),
		},
	}
}

// structs contructor function
func New(firstName, lastName, birthday string) (*User, error) {
	if firstName == "" || lastName == "" || birthday == "" {
		return nil, fmt.Errorf("firstname, lastname, birthday are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthday:  birthday,
		createdAt: time.Now(),
	}, nil
}
