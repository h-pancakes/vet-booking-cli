package main

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

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

	// this helps wipe traces of hashed and plain passwords from memory
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

	clear(password)
	u.passwordHash = ""

	return u, err
}

// Buffer between CLI validation layer with scanner and service layer. Core business logic here
func NewUser(firstName, lastName, phone, email string) (user, error) {
	firstName = strings.TrimSpace(firstName)
	if firstName == "" {
		return user{}, ErrNameTooShort
	}

	if len(firstName) > 20 {
		return user{}, ErrNameTooLong
	}

	for _, c := range firstName {
		if c >= 'A' && c <= 'Z' {
			continue
		}
		if c >= 'a' && c <= 'z' {
			continue
		}
		if c == ' ' {
			continue
		}
		if c == '-' {
			continue
		}
		return user{}, ErrNameContainsInvalidCharacters
	}

	lastName = strings.TrimSpace(lastName)
	if lastName == "" {
		return user{}, ErrNameTooShort
	}

	if len(lastName) > 20 {
		return user{}, ErrNameTooLong
	}

	for _, c := range lastName {
		if c >= 'A' && c <= 'Z' {
			continue
		}
		if c >= 'a' && c <= 'z' {
			continue
		}
		if c == ' ' {
			continue
		}
		if c == '-' {
			continue
		}
		return user{}, ErrNameContainsInvalidCharacters
	}

	phone = strings.ReplaceAll(phone, " ", "")

	if len(phone) < 10 {
		return user{}, ErrInvalidPhone
	}

	if len(phone) > 13 {
		return user{}, ErrInvalidPhone
	}

	for i, c := range phone {
		if i == 0 && c == '+' {
			continue
		}
		if c >= '0' && c <= '9' {
			continue
		}
		return user{}, ErrInvalidPhoneCharacters
	}

	email = strings.ReplaceAll(email, " ", "")

	email = strings.ToLower(email)

	if len(email) < 5 {
		return user{}, ErrInvalidEmail
	}

	if len(email) > 256 {
		return user{}, ErrInvalidEmail
	}

	count := 0
	for _, c := range email {
		if c == '@' {
			count++
		}
	}
	if count != 1 {
		return user{}, ErrInvalidEmail
	}

	parts := strings.SplitN(email, "@", 2)
	local := parts[0]
	domain := parts[1]

	if local == "" || domain == "" {
		return user{}, ErrInvalidEmail
	}

	for i, d := range local {
		if d >= 'A' && d <= 'Z' {
			continue
		}
		if d >= 'a' && d <= 'z' {
			continue
		}
		if d == '.' || d == '_' || d == '-' || d == '+' {
			continue
		}
		if d >= '0' && d <= '9' {
			continue
		}
		if i == 0 && d == '.' {
			return user{}, ErrInvalidEmail
		}
		if i == len(local)-1 && d == '.' {
			return user{}, ErrInvalidEmail
		}
		if i > 0 && local[i-1] == '.' && d == '.' {
			return user{}, ErrInvalidEmail
		}
		return user{}, ErrInvalidEmail
	}

	for i, e := range domain {
		if e >= 'A' && e <= 'Z' {
			continue
		}
		if e >= 'a' && e <= 'z' {
			continue
		}
		if e == '.' || e == '-' {
			continue
		}
		if i == 0 && e == '.' {
			return user{}, ErrInvalidEmail
		}
		if i == len(domain)-1 && e == '.' {
			return user{}, ErrInvalidEmail
		}
		if i > 0 && domain[i-1] == '.' && e == '.' {
			return user{}, ErrInvalidEmail
		}
		return user{}, ErrInvalidEmail
	}

	return user{firstName: firstName, lastName: lastName, phone: phone, email: email}, nil
}
