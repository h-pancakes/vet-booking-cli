package main

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo UserRepository
}

// this is a constructor - for abstraction
func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

// vertical slice for user side ONLY
func (s *UserService) RegisterUser(firstName, lastName, phone, email string, password []byte) (user, error) {
	u, err := NewUser(firstName, lastName, phone, email)
	if err != nil {
		return user{}, err
	}

	_, err = s.repo.FindByEmail(u.email)
	switch {
	case err == nil:
		return user{}, ErrDuplicateEmail
	case err != ErrEmailNotFound:
		return user{}, fmt.Errorf("checking email: %w", err)
	}

	if len(password) < 8 {
		return user{}, errors.New("password must be at least 8 characters long")
	}

	if len(password) > 72 {
		return user{}, errors.New("password must not be more than 72 bytes")
	}

	hashedBytes, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return user{}, fmt.Errorf("secure password generation failed: %w", err)
	}

	u.passwordHash = string(hashedBytes)

	err = s.repo.CreateUser(u)
	if err != nil {
		return user{}, fmt.Errorf("database save failed: %w", err)
	}

	// wipe traces of password from memory
	clear(password)
	u.passwordHash = ""

	return u, nil
}

func (s *UserService) LoginUser(email string, password []byte) (user, error) {
	u, hashedPassword, err := s.repo.FindByEmailForLogin(email)

	if err != nil {
		return user{}, fmt.Errorf("database query failed: %w", err)
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, password)

	if err != nil {
		return user{}, ErrInvalidEmailOrPassword
	}

	// wipe traces of password from memory
	clear(password)
	u.passwordHash = ""

	return u, err
}
