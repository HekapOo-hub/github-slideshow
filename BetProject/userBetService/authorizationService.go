package main

import (
	"golang.org/x/crypto/bcrypt"
)

type ValidationError struct {
	Message string
	Status  int
}

func (e ValidationError) Error() string {
	return e.Message
}

type authorization interface {
	Authorize(name, password string) (string, int, error)
	CreateUser(*User) (string, error)
	validate(*User) error
}
type AuthorizationService struct {
	repository
}

func (a *AuthorizationService) Authorize(name, password string) (string, int, error) {
	//hash password
	u, err := a.repository.GetByName(name)

	if err != nil {
		return "", -1, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {

		return "", -1, err
	}
	return u.Role, u.UserId, nil
}
func (a *AuthorizationService) CreateUser(u *User) (string, error) {
	var err error
	if err = a.validate(u); err == nil {
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return "Hashing error", err
		}

		u.Password = string(hashedPass)
		str, err := a.repository.CreateUser(u)
		return str, err
	}
	return "Creation error", err
}
func (a *AuthorizationService) validate(u *User) error {
	if u.Name == "" || u.Password == "" {
		return &ValidationError{Message: "Missing fields\n"}
	}
	if _, err := a.repository.GetByName(u.Name); err == nil {

		return &ValidationError{Message: "User with this name already exists.Please change your name"}
	}

	return nil
}
