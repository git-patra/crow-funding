package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(data RegisterUserInput) (User, error)
	Login(data LoginInput) (User, error)
	IsEmailAvailable(data CheckEmailInput) (bool, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo,
	}
}

func (create *service) Register(data RegisterUserInput) (User, error) {
	user := User{
		Name:       data.Name,
		Email:      data.Email,
		Occupation: data.Occupation,
		Role:       "user",
	}

	passwdHas, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MinCost)

	if err != nil {
		return user, nil
	}

	user.PasswordHash = string(passwdHas)

	newUser, err := create.repo.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repo.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found!")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return user, errors.New("Email or Password is wrong!")
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repo.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, errors.New("Email already exist!")
}
