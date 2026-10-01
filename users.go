package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/term"
)

var ErrLoginIdMustBePositive = errors.New("login ID must be a positive number")
var ErrNameTooShort = errors.New("name must be at least 1 character")
var ErrNameTooLong = errors.New("name cannot be more than 20 characters")
var ErrNameContainsInvalidCharacters = errors.New("name can only contain A-Z, hyphens, and spaces")
var ErrInvalidPhone = errors.New("phone number must be between 10 and 13 digits")
var ErrInvalidPhoneCharacters = errors.New("phone number can only contain digits 0-9")
var ErrInvalidEmail = errors.New("must be a valid email address")
var ErrDuplicateEmail = errors.New("email already taken")
var ErrEmailNotFound = errors.New("email not found")
var ErrInvalidEmailOrPassword = errors.New("invalid email or password")

// for future use:
var emailRe = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
var phoneRe = regexp.MustCompile(`^\d{10,11,12,13}$`)

// getExistingUser is a special function that is called when the user selects option "2" in the main menu to indicate they are an existing user.
// The function prompts the user to enter their login ID to access their appointments saved on the database.
// The user's email is normalised, then subsequently validated, and the database is queried for a row with a matching ID.
// If there is a matching ID, that row's contents are fetched and placed in memory.
func getExistingUser(scanner *bufio.Scanner, db *sql.DB) (*user, error) {
	fmt.Println("Please enter your login ID:")
	fmt.Print("> ")

	scanner.Scan()
	email := strings.TrimSpace(scanner.Text())

	id, err := strconv.Atoi(email)
	if err != nil || id <= 0 {
		return nil, fmt.Errorf("login ID must be a positive number")
	}

	var u user

	err = db.QueryRow(
		`SELECT id, first_name, last_name, phone, email
		 FROM users
		 WHERE id = $1`,
		id,
	).Scan(
		&u.id,
		&u.firstName,
		&u.lastName,
		&u.phone,
		&u.email,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("no user found with that ID")
	}
	if err != nil {
		return nil, err
	}

	fmt.Println("Welcome,", u.firstName)

	return &u, nil
}

// getUserPhone is a helper function that prompts the user for their phone number and stores it.
// The stored number is normalised by removing unnecessary whitespace.
// The number is passed through multiple validation checks and is returned if it passes all checks.
// If validation fails, an error is returned.
func getUserPhone(scanner *bufio.Scanner) (string, error) {
	var email string

	fmt.Println("Please enter your mobile phone number: ")
	scanner.Scan()

	email = scanner.Text()

	email = strings.ReplaceAll(email, " ", "")

	if len(email) < 10 {
		return "", fmt.Errorf("phone number must be more than 10 characters")
	}

	if len(email) > 13 {
		return "", fmt.Errorf("phone number must be smaller than 13 characters")
	}

	for i, c := range email {
		if i == 0 && c == '+' {
			continue
		}
		if c >= '0' && c <= '9' {
			continue
		}
		return "", fmt.Errorf("phone number can only have digits 0-9, +, and must be between 10 and 13 characters")
	}
	return email, nil
}

// getUserEmail is a helper function that prompts the user for their email address and stores it.
// The stored email address is normalised by removing unnecessary whitespace.
// The email address is passed through multiple validation checks and is returned if it passes all checks.
// If validation fails, an error is returned.
func getUserEmail(scanner *bufio.Scanner) (string, error) {
	var email string

	fmt.Println("Please enter your email address: ")
	scanner.Scan()

	email = scanner.Text()

	email = strings.ReplaceAll(email, " ", "")

	email = strings.ToLower(email)

	if len(email) < 5 {
		return "", fmt.Errorf("minimum email length is 5 characters")
	}

	if len(email) > 256 {
		return "", fmt.Errorf("maximum email length is 256 characters")
	}

	count := 0
	for _, c := range email {
		if c == '@' {
			count++
		}
	}
	if count != 1 {
		return "", fmt.Errorf("email must contain one @ symbol")
	}

	parts := strings.SplitN(email, "@", 2)
	local := parts[0]
	domain := parts[1]

	if local == "" || domain == "" {
		return "", fmt.Errorf("email must have text before and after '@'")
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
			return "", fmt.Errorf("cannot begin or end email with '.'")
		}
		if i == len(local)-1 && d == '.' {
			return "", fmt.Errorf("cannot begin or end email with '.'")
		}
		if i > 0 && local[i-1] == '.' && d == '.' {
			return "", fmt.Errorf("cannot have consecutive dots in first part of email")
		}
		return "", fmt.Errorf("invalid email email")
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
			return "", fmt.Errorf("cannot begin or end email with '.'")
		}
		if i == len(domain)-1 && e == '.' {
			return "", fmt.Errorf("cannot begin or end email with '.'")
		}
		if i > 0 && domain[i-1] == '.' && e == '.' {
			return "", fmt.Errorf("cannot have consecutive dots in second part of email")
		}
		return "", fmt.Errorf("invalid email email")
	}
	return email, nil
}

func getPassword(promptText string, minLength int) ([]byte, error) {
	for {
		fmt.Print(promptText)
		password, err := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println()

		if err != nil {
			clear(password)
			return nil, fmt.Errorf("read password: %w", err)
		}

		if len(password) < minLength {
			clear(password)
			fmt.Printf("Password must be at least %d bytes.\n", minLength)
			continue
		}

		if len(password) > 72 {
			clear(password)
			fmt.Println("Password must not exceed 72 bytes.")
			continue
		}

		return password, nil
	}
}

func gatherUserInfo(scanner *bufio.Scanner) (string, string, string, string, []byte, error) {
	fmt.Println("Please enter your user information")

	var firstName, lastName, phone, email string
	var password []byte
	var err error

	for {
		firstName, err = getName(scanner, "Please type in your first name: ")
		if err == nil {
			break
		}
		fmt.Println("Error: ", err)
	}
	for {
		lastName, err = getName(scanner, "Please type in your last name: ")
		if err == nil {
			break
		}
		fmt.Println("Error: ", err)
	}

	for {
		phone, err = getUserPhone(scanner)
		if err == nil {
			break
		}
		fmt.Println("Error: ", err)
	}

	for {
		email, err = getUserEmail(scanner)
		if err == nil {
			break
		}
		fmt.Println("Error: ", err)
	}

	password, err = getPassword("Please enter a password: ", 8)
	if err != nil {
		return "", "", "", "", nil, err
	}

	return firstName, lastName, phone, email, password, nil
}

// String prints a summary of the user's details.
func (u user) String() string {
	var s string
	s = "-------------------------------------\n"
	s += "Owner details:\n"
	s += fmt.Sprintf("Name: %s\n", u.firstName)
	s += fmt.Sprintf("Surname: %s\n", u.lastName)
	s += fmt.Sprintf("Phone number: %s\n", u.phone)
	s += fmt.Sprintf("Email address: %s\n", u.email)
	s += "-------------------------------------\n"

	return s
}

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
