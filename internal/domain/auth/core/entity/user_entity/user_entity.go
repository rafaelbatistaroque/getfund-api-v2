package user_entity

import (
	"time"
)

type User struct {
	firstName string
	lastName  string
	username  string
	password  string
	isActive  bool
	isAdmin   bool
	createdAt time.Time
	updatedAt time.Time
}

func New(firstName, lastName, username, password string) *User {
	return &User{
		firstName: getValidValue(firstName),
		lastName:  getValidValue(lastName),
		username:  getValidValue(username),
		password:  getValidValue(password),
		isActive:  true,
		isAdmin:   false,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func getValidValue(value string) string {
	if value == "" {
		panic("error on create user entity")
	}

	return value
}

func (u *User) GetFirstName() string {
	return u.firstName
}

func (u *User) GetLastName() string {
	return u.lastName
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetIsActive() bool {
	return u.isActive
}

func (u *User) GetIsAdmin() bool {
	return u.isAdmin
}

func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.updatedAt
}
