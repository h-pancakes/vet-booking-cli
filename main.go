package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

// user is a struct that holds information about the user of the booking service.
// It contains contact details persisted to the database.
type user struct {
	firstName string
	lastName  string
	phone     string
	email     string
}

// pet is a struct that holds information about a pet that the user is booking an appointment for.
// This information is stored in the appointments table in the database.
type pet struct {
	name       string
	species    string
	age        int
	weightKg   float64
	vaccinated bool
}

// appointment is a struct that holds all information related to an appointment booked by the user.
// This information is stored in the appointments table in the database.
type appointment struct {
	appointmentType string
	pet             pet
	vet             string
	dateTime        time.Time
}

// allowedSpecies is a list that holds the options for choosing the pet's species for the appointment.
var allowedSpecies = []string{
	"Dog",
	"Cat",
}

// allowedAppointmentTypes is a list that holds the types of appointments available to the user.
var allowedAppointmentTypes = []string{
	"Grooming",
	"Vaccination",
	"Surgical",
	"Bath",
	"Dental",
}

// allowedVets is a list that holds the veterinarians that are available to the user.
var allowedVets = []string{
	"Dr Smith",
	"Dr Jones",
	"Dr Patel",
	"Dr Brown",
}

// getUserFirstName is a helper function that prompts the user for their first name and then stores it.
// The stored name is then normalised by removing unnecessary whitespace.
// The name is passed through multiple validation checks and returned, if it passes all checks.
// If validation fails, an error is returned.
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

// mainMenu is a function displays a menu screen to the user with 3 options.
// The option that the user selects is normalised and then passed to main().
func mainMenu(scanner *bufio.Scanner) string {
	fmt.Println("1. Create new appointment")
	fmt.Println("2. View appointments")
	fmt.Println("3. Exit")
	fmt.Print("> ")

	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// petCounter is a function that prompts the user to enter the number of pets they wish to book an appointment for.
// The input is converted into an integer type and then validated.
// If the input is invalid, an error is returned.
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

// getName is a helper function that prompts the user for their pet's name and stores it.
// The stored name is normalised by removing unnecessary whitespace.
// The name is passed through multiple validation checks and is returned if it passes all checks.
// If validation fails, an error is returned.
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

// getSpecies is a helper function that prompts the user to provide their pet's species and lists available options using the "allowedSpecies" list.
// The input is stored and normalised.
// If the input is not listed in "allowedSpecies", the user is prompted again.
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

// getAge is a helper function that prompts the user for their pet's age and stores it.
// The stored age is converted to an integer type.
// The name is passed through a validation check.
// If validation fails, an error is returned.
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

// getWeightKg is a helper function that prompts the user for their pet's weight in kilograms and stores it.
// The stored weight is converted to a float64 type.
// The weight is passed through a validation check.
// If validation fails, an error is returned.
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

// getVaccinationStatus is a helper function that prompts the user to clarify whether their pet is vaccinated or not.
// The function takes the user input in the form of a (y/n) and stores it in an input variable.
// The variable is passed through a switch statement that either stores a boolean value or returns an error in the case of an invalid input.
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

// getAppointmentType is a helper function that prompts the user to choose an appointment type and lists available options using the "allowedAppointmentTypes" list.
// The input is stored and normalised.
// If the input is not listed in "allowedAppointmentTypes", the user is prompted again.
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

// getVet is a helper function that prompts the user to choose a preferred vet for their appointment and lists available options using the "allowedVets" list.
// The input is stored and normalised.
// If the input is not listed in "allowedAppointmentTypes", the user is prompted again.
func getVet(scanner *bufio.Scanner, i int) (string, error) {
	var input string

	fmt.Println("Please choose preferred vet for appointment", i+1)
	fmt.Println("Available vets:", allowedVets)

	scanner.Scan()
	input = strings.TrimSpace(scanner.Text())

	for _, v := range allowedVets {
		if v == input {
			return v, nil
		}
	}

	return "", fmt.Errorf("please choose a valid vet")
}

// getPreferredDateTime is a helper function that allows the user to enter a preferred date and time for their appointment.
// The user is prompted for a date and time in a specified format.
// The input is stored and normalised.
// The input is parsed and converted into a date and time format.
// The input is then validated and an error is displayed if it doesn't pass the validation checks.
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

// bookAppointments calls the helper functions repeatedly until a valid input is received from the user for all fields. This procedure is iterated for each appointment the user filled in details for.
// If an error is received for a helper function, bookAppointments calls the function again, and the user is prompted for a valid input.
// If a valid input is received for a helper function, bookAppointments will pass the valid input to the corresponding field in the newly initialised "appointment" objects.
// The appointment objects are stored in a list to accommodate multiple appointments.
// Once all fields in "appointment" are filled, bookAppointments returns the list of "appointment" objects.
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

// summaryString prints a summary of each appointment's details.
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
	s += fmt.Sprintf("Vet: %s\n", a.vet)
	s += fmt.Sprintf("Appointment Date & Time: %s\n", a.dateTime.Format("Monday, 02 Jan 2006 at 15:04"))
	s += "-------------------------------------\n"

	return s
}

// ownerSummaryString prints a summary of the user's details.
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
	var currentUser *user
	var appointments []appointment

	scanner := bufio.NewScanner(os.Stdin)

	db, err := sql.Open(
		"postgres",
		"user=vetapp password=vetpass dbname=vet_booking sslmode=disable",
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	u := gatherUserInfo(scanner)
	currentUser = &u

	var userID int

	err = db.QueryRow(
		`INSERT INTO users (first_name, last_name, phone, email)
		VALUES ($1, $2, $3, $4)
		RETURNING id`,
		u.firstName,
		u.lastName,
		u.phone,
		u.email,
	).Scan(&userID)
	if err != nil {
		panic(err)
	}

	for {
		userChoice := mainMenu(scanner)

		switch userChoice {
		case "1":
			var petCount int
			for {
				count, err := petCounter(scanner)
				if err == nil {
					petCount = count
					break
				}
				fmt.Println("Error:", err)
			}

			newAppointments := bookAppointments(scanner, petCount)
			appointments = append(appointments, newAppointments...)

			for _, a := range newAppointments {
				_, err := db.Exec(
					`INSERT INTO appointments (
					user_id,
					pet_name,
					pet_species,
					pet_age,
					pet_weight,
					vaccinated,
					appointment_type,
					vet_name,
					appointment_time
					) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
					userID,
					a.pet.name,
					a.pet.species,
					a.pet.age,
					a.pet.weightKg,
					a.pet.vaccinated,
					a.appointmentType,
					a.vet,
					a.dateTime,
				)
				if err != nil {
					panic(err)
				}
			}

		case "2":
			if len(appointments) == 0 {
				fmt.Println("No appointments yet.")
				continue
			}

			fmt.Println(currentUser.ownerSummaryString())
			for i, a := range appointments {
				fmt.Println(a.summaryString(i + 1))
			}
		case "3":
			fmt.Println("Thank you for using our booking service!")
			return

		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
