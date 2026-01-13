package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// user holds information about the user of the booking service
type user struct {
	firstName string
	lastName  string
	phone     string
	email     string
}

// pet holds information about a pet
type pet struct {
	name       string
	species    string
	age        int
	weightKg   float64
	vaccinated bool
}

// appointment holds all information related to an appointment
type appointment struct {
	appointmentType string
	pet             pet
	vet             vet
	dateTime        time.Time
}

// vet holds information about vets
type vet struct {
	name string
}

// allowedSpecies holds the species options available to the user
var allowedSpecies = []string{
	"Dog",
	"Cat",
}

// allowedAppointmentTypes holds the appointment type options available
var allowedAppointmentTypes = []string{
	"Grooming",
	"Vaccination",
	"Surgical",
	"Bath",
	"Dental",
}

// allowedVets holds the veterinarians available to the user
var allowedVets = []string{
	"Dr Smith",
	"Dr Jones",
	"Dr Patel",
	"Dr Brown",
}

// getUserFirstName prompts the user for their first name and validates it
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

// getUserLastName prompts the user for their surname and validates it
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

// getUserPhone prompts the user for their telephone number and validates it
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

// getUserEmail prompts the user for their email address and validates it
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

// gatherUserInfo collects validated inputs for user and creates a user object
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

// petCounter prompts the user to enter the number of pets they wish to book an appointment for
func petCounter(scanner *bufio.Scanner) (int, error) {
	var petCount int

	fmt.Println("Please enter how many pets you are booking appointments for: ")
	scanner.Scan()
	petCount, _ = strconv.Atoi(scanner.Text())

	if petCount > 0 && petCount <= 20 {
		return petCount, nil
	} else if petCount > 20 {
		return 0, fmt.Errorf("input exceeds limit of 20")
	} else {
		return 0, fmt.Errorf("invalid input")
	}
}

// getName prompts the user for a pet's name and validates it
func getName(scanner *bufio.Scanner, i int) (string, error) {
	var input string

	fmt.Println("Please enter pet", i+1, "name: ")
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

// getSpecies asks for a pet's species and ensures it matches an allowed option
func getSpecies(scanner *bufio.Scanner, i int) (string, error) {
	var input string

	fmt.Println("Please enter pet", i+1, "species: ")
	fmt.Println("Accepted species:", allowedSpecies)

	scanner.Scan()
	input = scanner.Text()
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	input = strings.Title(input)

	valid := false

	for _, species := range allowedSpecies {
		if species == input {
			valid = true
			break
		}
	}

	if !valid {
		return "", fmt.Errorf("please enter an accepted species")
	}
	return input, nil
}

// getAge prompts the user for a pet's age and validates the range
func getAge(scanner *bufio.Scanner, i int) (int, error) {
	var input int

	fmt.Println("Please enter pet", i+1, "age: ")
	scanner.Scan()
	input, _ = strconv.Atoi(scanner.Text())

	if input < 0 || input > 30 {
		return 0, fmt.Errorf("age must be between 0 and 30 years")
	}
	return input, nil
}

// getWeightKg prompts for the pet's weight in kilograms and validates the range
func getWeightKg(scanner *bufio.Scanner, i int) (float64, error) {
	var input float64

	fmt.Println("Please enter pet", i+1, "weight (Kg): ")
	scanner.Scan()
	input, _ = strconv.ParseFloat(scanner.Text(), 64)

	if input < 1 || input > 120 {
		return 0, fmt.Errorf("weight must be between 1 and 120kg")
	}
	return input, nil
}

// getVaccinationStatus prompts for and validates the pet's vaccination status
func getVaccinationStatus(scanner *bufio.Scanner, i int) (bool, error) {
	var input string

	fmt.Println("Is pet", i+1, "vaccinated? (y/n): ")
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

// getAppointmentType displays available appointment types and ensures the user picks an accepted type
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

// getVet displays available vets and ensures the user chooses a preferred vet
func getVet(scanner *bufio.Scanner, i int) (vet, error) {
	var input string

	fmt.Println("Please choose preferred vet for appointment", i+1)
	fmt.Println("Available vets:", allowedVets)

	scanner.Scan()
	input = strings.TrimSpace(scanner.Text())

	for _, v := range allowedVets {
		if v == input {
			return vet{name: v}, nil
		}
	}

	return vet{}, fmt.Errorf("please choose a valid vet")
}

func getPreferredDateTime(scanner *bufio.Scanner, i int) (time.Time, error) {
	fmt.Println("Please enter preferred date and time for appointment", i+1)
	fmt.Println("Format: YYYY-MM-DD HH:MM (24-hour time)")
	fmt.Println("Example: 2026-01-13 12:30")

	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	layout := "2006-01-02 15:04"
	t, err := time.Parse(layout, input)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date/time format")
	}

	if t.Before(time.Now()) {
		return time.Time{}, fmt.Errorf("Appointment cannot be in the past")
	}

	return t, nil
}

// bookAppointment collects validated input for each pet and returns a slice of pets
func bookAppointments(scanner *bufio.Scanner, petCount int) []appointment {
	appointments := make([]appointment, 0, petCount)

	for i := 0; i < petCount; i++ {

		var d pet

		for {
			name, err := getName(scanner, i)
			if err == nil {
				d.name = name
				break
			}
			fmt.Println("Error:", err)
		}

		for {
			breed, err := getSpecies(scanner, i)
			if err == nil {
				d.species = breed
				break
			}
			fmt.Println("Error:", err)
		}

		for {
			age, err := getAge(scanner, i)
			if err == nil {
				d.age = age
				break
			}
			fmt.Println("Error:", err)
		}

		for {
			weightKg, err := getWeightKg(scanner, i)
			if err == nil {
				d.weightKg = weightKg
				break
			}
			fmt.Println("Error:", err)
		}

		for {
			vaccinated, err := getVaccinationStatus(scanner, i)
			if err == nil {
				d.vaccinated = vaccinated
				break
			}
			fmt.Println("Error:", err)
		}

		var a appointment

		for {
			appointmentType, err := getAppointmentType(scanner, i)
			if err == nil {
				a.appointmentType = appointmentType
				break
			}
			fmt.Println("Error:", err)
		}

		for {
			v, err := getVet(scanner, i)
			if err == nil {
				a.vet = v
				break
			}
			fmt.Println("Error:", err)
		}

		for {
			dt, err := getPreferredDateTime(scanner, i)
			if err == nil {
				a.dateTime = dt
				break
			}
			fmt.Println("Error:", err)
		}

		a.pet = d
		appointments = append(appointments, a)
	}

	return appointments
}

// summaryString prints a summary of the pet and appointment details
func (a *appointment) summaryString(i int) string {
	var s string
	s = "-------------------------------------\n"
	s += fmt.Sprintf("Summary for Appointment %d\n", i)
	s += fmt.Sprintf("Pet Name: %s\n", a.pet.name)
	s += fmt.Sprintf("Species: %s\n", a.pet.species)
	s += fmt.Sprintf("Age: %d\n", a.pet.age)
	s += fmt.Sprintf("Weight (kg): %.2f\n", a.pet.weightKg)
	s += fmt.Sprintf("Vaccinated?: %t\n", a.pet.vaccinated)
	s += fmt.Sprintf("Appointment Type: %s\n", a.appointmentType)
	s += fmt.Sprintf("Vet: %s\n", a.vet.name)
	s += fmt.Sprintf("Appointment Date & Time: %s\n", a.dateTime.Format("Monday, 02 Jan 2006 at 15:04"))
	s += "-------------------------------------\n"

	return s
}

// ownerSummaryString prints a summary of the user's details
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

	var petCount int

	for {
		count, err := petCounter(scanner)
		if err == nil {
			petCount = count
			break
		}
		fmt.Println("Error: ", err)
	}

	appointments := bookAppointments(scanner, petCount)

	fmt.Println(user.ownerSummaryString())

	for i, d := range appointments {
		fmt.Println(d.summaryString(i + 1))
	}

	fmt.Println("Thank you for using our booking service!")
}
