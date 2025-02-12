package activate_user_entity

import (
	"slices"
	"time"
)

type ActivationUser struct {
	firstName         string
	lastName          string
	email             string
	gender            string
	password          string
	countryId         int
	userCategoryId    int
	mainSocialNetwork string
	registeredUrl     string
	registeredAt      time.Time
	isActive          bool
	isAdmin           bool
}

func New(firstName, lastName, email, gender, password, mainSocialNetwork, registeredUrl string, countryId, userCategoryId int) *ActivationUser {
	return &ActivationUser{
		firstName:         getValidValue(firstName),
		lastName:          getValidValue(lastName),
		email:             getValidValue(email),
		gender:            validateGender(gender),
		password:          getValidValue(password),
		countryId:         countryId,
		userCategoryId:    userCategoryId,
		mainSocialNetwork: getValidValue(mainSocialNetwork),
		registeredUrl:     getValidValue(registeredUrl),
		registeredAt:      time.Now(),
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

func (u *ActivationUser) GetFirstName() string {
	return u.firstName
}

func (u *ActivationUser) GetLastName() string {
	return u.lastName
}

func (u *ActivationUser) GetEmail() string {
	return u.email
}

func (u *ActivationUser) GetGender() string {
	return u.gender
}

func (u *ActivationUser) GetPassword() string {
	return u.password
}

func (u *ActivationUser) GetCountryId() int {
	return u.countryId
}

func (u *ActivationUser) GetUserCategoryId() int {
	return u.userCategoryId
}

func (u *ActivationUser) GetMainSocialNetwork() string {
	return u.mainSocialNetwork
}

func (u *ActivationUser) GetRegisteredUrl() string {
	return u.registeredUrl
}

func (u *ActivationUser) GetRegisteredAt() int64 {
	return u.registeredAt.Unix()
}

func (u *ActivationUser) GetIsActive() bool {
	return u.isActive
}

func (u *ActivationUser) GetIsAdmin() bool {
	return u.isAdmin
}
