package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type user struct {
	firstName string
	lastName  string
	phone     string
	email     string
}

// dog holds information about a pet being booked for an appointment.
type dog struct {
	name            string
	breed           string
	age             int
	weightKg        float64
	vaccinated      bool
	appointmentType string
}

var allowedBreeds = []string{
	"Labrador",
	"Poodle",
	"Beagle",
	"Bulldog",
	"Dachshund",
}

var allowedAppointmentTypes = []string{
	"Grooming",
	"Vaccination",
	"Surgical",
	"Bath",
	"Dental",
}

func getUserFirstName(scanner *bufio.Scanner) (string, error) {
	var input string

	fmt.Println("Welcome to our booking service!")
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
	return input, nil
}

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
	return input, nil
}

func getUserPhone(scanner *bufio.Scanner) (string, error) {
	var input string

	fmt.Println("Please enter your mobile phone number: ")
	scanner.Scan()

	input = scanner.Text()

	input = strings.ReplaceAll(input, " ", "")

	if len(input) < 10 {
		return "", fmt.Errorf("invalid phone number")
	}

	if len(input) > 13 {
		return "", fmt.Errorf("invalid phone number")
	}

	for i, c := range input {
		if i == 0 && c == '+' {
			continue
		}
		if c >= '0' && c <= '9' {
			continue
		}
		return "", fmt.Errorf("invalid phone number")
	}
	return input, nil
}

func getUserEmail(scanner *bufio.Scanner) (string, error) {
	var input string

	fmt.Println("Please enter your email address: ")
	scanner.Scan()

	input = scanner.Text()

	input = strings.ReplaceAll(input, " ", "")

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

func gatherUserInfo(scanner *bufio.Scanner) user {
	var user user

	for {
		firstName, err := getUserFirstName(scanner)
		if err == nil {
			user.firstName = firstName
			break
		}
		fmt.Println("Error: ", err)
	}
	for {
		lastName, err := getUserLastName(scanner)
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

// dogCounter prompts the user to enter the number of dogs they wish to book an appointment for.
func dogCounter(scanner *bufio.Scanner) (int, error) {
	var dogCount int

	fmt.Println("Please enter how many dogs you are booking appointments for: ")
	scanner.Scan()
	dogCount, _ = strconv.Atoi(scanner.Text())

	if dogCount > 0 && dogCount <= 20 {
		return dogCount, nil
	} else if dogCount > 20 {
		return 0, fmt.Errorf("input exceeds limit of 20")
	} else {
		return 0, fmt.Errorf("invalid input")
	}
}

// getName prompts the user for a dog's name and validates it.
func getName(scanner *bufio.Scanner, i int) (string, error) {
	var input string

	fmt.Println("Please enter dog", i+1, "name: ")
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
		return "", fmt.Errorf("name can only contain characters A-Z and spaces")
	}
	return input, nil
}

// getBreed asks for a dog's breed and ensures it matches an allowed option.
func getBreed(scanner *bufio.Scanner, i int) (string, error) {
	var input string

	fmt.Println("Please enter dog", i+1, "breed: ")
	fmt.Println("Accepted Breeds:", allowedBreeds)

	scanner.Scan()
	input = scanner.Text()
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	input = strings.Title(input)

	valid := false

	for _, breed := range allowedBreeds {
		if breed == input {
			valid = true
			break
		}
	}

	if !valid {
		return "", fmt.Errorf("please enter an accepted breed")
	}
	return input, nil
}

// getAge prompts the user for a dog's age and validates the range.
func getAge(scanner *bufio.Scanner, i int) (int, error) {
	var input int

	fmt.Println("Please enter dog", i+1, "age: ")
	scanner.Scan()
	input, _ = strconv.Atoi(scanner.Text())

	if input < 0 || input > 30 {
		return 0, fmt.Errorf("age must be between 0 and 30 years")
	}
	return input, nil
}

// getWeightKg prompts for the dog's weight in kilograms and validates the range.
func getWeightKg(scanner *bufio.Scanner, i int) (float64, error) {
	var input float64

	fmt.Println("Please enter dog", i+1, "weight (Kg): ")
	scanner.Scan()
	input, _ = strconv.ParseFloat(scanner.Text(), 64)

	if input < 1 || input > 120 {
		return 0, fmt.Errorf("weight must be between 1 and 120kg")
	}
	return input, nil
}

// getVaccinationStatus prompts for and validates the dog's vaccination status.
func getVaccinationStatus(scanner *bufio.Scanner, i int) (bool, error) {
	var input string

	fmt.Println("Is dog", i+1, "vaccinated? (y/n): ")
	scanner.Scan()
	input = scanner.Text()

	switch input {
	case "y", "Y":
		return true, nil
	case "n", "N":
		return false, nil
	default:
		return false, fmt.Errorf("invalid input")
	}
}

// getAppointmentType displays available appointment types and ensures the user picks an accepted type.
func getAppointmentType(scanner *bufio.Scanner, i int) (string, error) {
	var input string

	fmt.Println("Please enter your requested appointment type for", i+1)
	fmt.Println("Accepted types: ", allowedAppointmentTypes)

	scanner.Scan()
	input = scanner.Text()
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	input = strings.Title(input)

	valid := false

	for _, appointmentType := range allowedAppointmentTypes {
		if appointmentType == input {
			valid = true
			break
		}
	}

	if !valid {
		return "", fmt.Errorf("please enter an accepted appointment type")
	}
	return input, nil
}

// bookAppointment collects validated input for each dog and returns a slice of dogs.
func bookAppointment(scanner *bufio.Scanner, dogCount int) []dog {
	dogs := make([]dog, dogCount)

	for i := 0; i < dogCount; i++ {

		for {
			name, err := getName(scanner, i)
			if err == nil {
				dogs[i].name = name
				break
			}
			fmt.Println("Error: ", err)
		}

		for {
			breed, err := getBreed(scanner, i)
			if err == nil {
				dogs[i].breed = breed
				break
			}
			fmt.Println("Error: ", err)
		}

		for {
			age, err := getAge(scanner, i)
			if err == nil {
				dogs[i].age = age
				break
			}
			fmt.Println("Error: ", err)
		}

		for {
			weightKg, err := getWeightKg(scanner, i)
			if err == nil {
				dogs[i].weightKg = weightKg
				break
			}
			fmt.Println("Error: ", err)
		}

		for {
			vaccinationStatus, err := getVaccinationStatus(scanner, i)
			if err == nil {
				dogs[i].vaccinated = vaccinationStatus
				break
			}
			fmt.Println("Error: ", err)
		}

		for {
			appointmentType, err := getAppointmentType(scanner, i)
			if err == nil {
				dogs[i].appointmentType = appointmentType
				break
			}
			fmt.Println("Error: ", err)
		}
	}
	return dogs
}

// summaryString prints a summary of each dog's details
func (d *dog) summaryString(i int) string {
	var s string
	s = "-------------------------------------\n"
	s += fmt.Sprintf("Summary for Dog %d\n", i)
	s += fmt.Sprintf("Name: %s\n", d.name)
	s += fmt.Sprintf("Breed: %s\n", d.breed)
	s += fmt.Sprintf("Age: %d\n", d.age)
	s += fmt.Sprintf("Weight (kg): %.2f\n", d.weightKg)
	s += fmt.Sprintf("Vaccinated?: %t\n", d.vaccinated)
	s += fmt.Sprintf("Appointment Type: %s\n", d.appointmentType)
	s += "-------------------------------------\n"

	return s

}

func (u *user) ownerSummaryString() string {
	var s string
	s = "-------------------------------------\n"
	s += "Summary for owner:\n"
	s += fmt.Sprintf("Name: %s\n", u.firstName)
	s += fmt.Sprintf("Surname: %s\n", u.lastName)
	s += fmt.Sprintf("Phone number: %s\n", u.phone)
	s += fmt.Sprintf("Email address: %s\n", u.email)
	s += "-------------------------------------\n"

	return s

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	user := gatherUserInfo(scanner)

	var dogCount int

	for {
		count, err := dogCounter(scanner)
		if err == nil {
			dogCount = count
			break
		}
		fmt.Println("Error: ", err)
	}

	dogs := bookAppointment(scanner, dogCount)

	fmt.Println(user.ownerSummaryString())

	for i, d := range dogs {
		fmt.Println(d.summaryString(i + 1))
	}

	fmt.Println("Thank you for using our booking service!")
}
