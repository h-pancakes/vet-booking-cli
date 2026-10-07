package main

import (
	"bufio"
	"fmt"
	"strings"
)

// getMainMenuChoice is a function that displays a menu screen to the user with 3 options.
// The option that the user selects is normalised and then passed to main().
func getMainMenuChoice(scanner *bufio.Scanner) string {
	fmt.Println("1. New user")
	fmt.Println("2. Existing user")
	fmt.Println("3. Exit")
	fmt.Print("> ")

	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// Shared by both user and pet name captures
func getName(scanner *bufio.Scanner, prompt string) (string, error) {
	var input string

	fmt.Println(prompt)
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
	// TODO: Change this to use new way
	input = strings.Title(input)
	return input, nil
}
