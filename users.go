package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

// getExistingUser is a special function that is called when the user selects option "2" in the main menu to indicate they are an existing user.
// The function prompts the user to enter their login ID to access their appointments saved on the database.
// The user's input is normalised, then subsequently validated, and the database is queried for a row with a matching ID.
// If there is a matching ID, that row's contents are fetched and placed in memory.
func getExistingUser(scanner *bufio.Scanner, db *sql.DB) (*user, error) {
	fmt.Println("Please enter your login ID:")
	fmt.Print("> ")

	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	id, err := strconv.Atoi(input)
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

// getUserFirstName is a helper function that prompts the user for their first name and then stores it.
// The stored name is then normalised by removing unnecessary whitespace.
// The name is passed through multiple validation checks and returned, if it passes all checks.
// If validation fails, an error is returned.
func getUserFirstName(scanner *bufio.Scanner) (string, error) {
	var input string

	fmt.Println("Please enter your first name: ")
	scanner.Scan()

	input = scanner.Text()

	input = strings.TrimSpace(input)

	trimmedInput := strings.ReplaceAll(input, " ", "")

	if len(trimmedInput) < 1 {
		return "", fmt.Errorf("name must be at least 1 character")
	}

	if len(trimmedInput) > 20 {
		return "", fmt.Errorf("character limit is 20 characters")
	}

	for _, c := range input {
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
		return "", fmt.Errorf("name can only contain A-Z, hyphens, and spaces")
	}
	// TODO: Change this to use supported method
	input = strings.Title(input)
	return input, nil
}

// getUserLastName is a helper function that prompts the user for their first name and then stores it.
// The stored name is then normalised by removing unnecessary whitespace.
// The name is passed through multiple validation checks and is returned if it passes all checks.
// If validation fails, an error is returned.
func getUserLastName(scanner *bufio.Scanner) (string, error) {
	var input string

	fmt.Println("Please enter your last name: ")
	scanner.Scan()

	input = scanner.Text()

	input = strings.TrimSpace(input)

	trimmedInput := strings.ReplaceAll(input, " ", "")

	if len(trimmedInput) < 1 {
		return "", fmt.Errorf("name must be at least 1 character")
	}

	if len(trimmedInput) > 20 {
		return "", fmt.Errorf("character limit is 20 characters")
	}

	for _, c := range input {
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
		return "", fmt.Errorf("name can only contain A-Z, hyphens, and spaces")
	}

	input = strings.Title(input)
	return input, nil
}

// getUserPhone is a helper function that prompts the user for their phone number and stores it.
// The stored number is normalised by removing unnecessary whitespace.
// The number is passed through multiple validation checks and is returned if it passes all checks.
// If validation fails, an error is returned.
func getUserPhone(scanner *bufio.Scanner) (string, error) {
	var input string

	fmt.Println("Please enter your mobile phone number: ")
	scanner.Scan()

	input = scanner.Text()

	input = strings.ReplaceAll(input, " ", "")

	if len(input) < 10 {
		return "", fmt.Errorf("phone number must be more than 10 characters")
	}

	if len(input) > 13 {
		return "", fmt.Errorf("phone number must be smaller than 13 characters")
	}

	for i, c := range input {
		if i == 0 && c == '+' {
			continue
		}
		if c >= '0' && c <= '9' {
			continue
		}
		return "", fmt.Errorf("phone number can only have digits 0-9, +, and must be between 10 and 13 characters")
	}
	return input, nil
}

// getUserEmail is a helper function that prompts the user for their email address and stores it.
// The stored email address is normalised by removing unnecessary whitespace.
// The email address is passed through multiple validation checks and is returned if it passes all checks.
// If validation fails, an error is returned.
func getUserEmail(scanner *bufio.Scanner) (string, error) {
	var input string

	fmt.Println("Please enter your email address: ")
	scanner.Scan()

	input = scanner.Text()

	input = strings.ReplaceAll(input, " ", "")

	input = strings.ToLower(input)

	if len(input) < 5 {
		return "", fmt.Errorf("minimum email length is 5 characters")
	}

	if len(input) > 256 {
		return "", fmt.Errorf("maximum email length is 256 characters")
	}

	count := 0
	for _, c := range input {
		if c == '@' {
			count++
		}
	}
	if count != 1 {
		return "", fmt.Errorf("email must contain one @ symbol")
	}

	parts := strings.SplitN(input, "@", 2)
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
		return "", fmt.Errorf("invalid email input")
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
		return "", fmt.Errorf("invalid email input")
	}
	return input, nil
}

// gatherUserInfo calls the helper functions repeatedly until a valid input is received from the user for all fields.
// If an error is received for a helper function, gatherUserInfo calls the function again, and the user is prompted for a valid input.
// If a valid input is received for a helper function, gatherUserInfo will pass the valid input to the corresponding field in the newly initialised "user" object.
// Once all fields in "user" are filled, gatherUserInfo returns the "user" object.
func gatherUserInfo(scanner *bufio.Scanner) user {
	var user user
	fmt.Println("Please enter your user information")

	for {
		firstName, err := getName(scanner, "Please type in your first name: ")
		if err == nil {
			user.firstName = firstName
			break
		}
		fmt.Println("Error: ", err)
	}
	for {
		lastName, err := getName(scanner, "Please type in your last name: ")
		if err == nil {
			user.lastName = lastName
			break
		}
		fmt.Println("Error: ", err)
	}

	for {
		phone, err := getUserPhone(scanner)
		if err == nil {
			user.phone = phone
			break
		}
		fmt.Println("Error: ", err)
	}

	for {
		email, err := getUserEmail(scanner)
		if err == nil {
			user.email = email
			break
		}
		fmt.Println("Error: ", err)
	}
	return user
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

func createNewUser(userToCreate user, db *sql.DB) bool {
	var err error
	var createdUserID int
	err = db.QueryRow(
		`INSERT INTO users (first_name, last_name, phone, email)
				 VALUES ($1, $2, $3, $4)
				 RETURNING id`,
		userToCreate.firstName,
		userToCreate.lastName,
		userToCreate.phone,
		userToCreate.email,
	).Scan(&createdUserID)
	if err != nil {
		fmt.Println("Error saving user to database:", err)
		return false
	}

	fmt.Println("Your login ID is:", createdUserID)
	fmt.Println("IMPORTANT: Save your login ID as you will need it to log in!")
	return true
}
