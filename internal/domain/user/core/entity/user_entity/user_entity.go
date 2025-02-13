package user_entity

import (
	"slices"
)

type User struct {
	firstName         string
	lastName          string
	email             string
	gender            string
	password          string
	countryId         int
	userCategoryId    int
	mainSocialNetwork string
	registeredUrl     string
	isActive          bool
	isAdmin           bool
}

func New(firstName, lastName, email, gender, password, mainSocialNetwork, registeredUrl string, countryId, userCategoryId int) *User {
	return &User{
		firstName:         getValidValue(firstName),
		lastName:          getValidValue(lastName),
		email:             getValidValue(email),
		gender:            validateGender(gender),
		password:          getValidValue(password),
		countryId:         countryId,
		userCategoryId:    userCategoryId,
		mainSocialNetwork: getValidValue(mainSocialNetwork),
		registeredUrl:     getValidValue(registeredUrl),
		isActive:          true,
		isAdmin:           false,
	}
}

func getValidValue(value string) string {
	if value == "" {
		panic("error on create user entity")
	}

	return value
}

func validateGender(gender string) string {
	if gender == "" || !slices.Contains([]string{"f", "m", "u", "nb"}, gender) {
		panic("gender is invalid")
	}

	return gender
}

func (u *User) GetFirstName() string {
	return u.firstName
}

func (u *User) GetLastName() string {
	return u.lastName
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetGender() string {
	return u.gender
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetCountryId() int {
	return u.countryId
}

func (u *User) GetUserCategoryId() int {
	return u.userCategoryId
}

func (u *User) GetMainSocialNetwork() string {
	return u.mainSocialNetwork
}

func (u *User) GetRegisteredUrl() string {
	return u.registeredUrl
}

func (u *User) GetIsActive() bool {
	return u.isActive
}

func (u *User) GetIsAdmin() bool {
	return u.isAdmin
}
