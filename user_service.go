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
func (s *UserService) RegisterUser(firstName, lastName, phone, email, password string) (user, error) {
	u, err := NewUser(firstName, lastName, phone, email)
	if err != nil {
		return user{}, err
	}

	// finds out if email is already taken
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

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return user{}, fmt.Errorf("secure password generation failed: %w", err)
	} // poop

	u.passwordHash = string(hashedBytes)

	err = s.repo.CreateUser(u)
	if err != nil {
		return user{}, fmt.Errorf("database save failed: %w", err)
	}

	return u, nil
}
