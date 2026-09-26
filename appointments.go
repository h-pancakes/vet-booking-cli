package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// getAppointmentsByUserID is a special function that is called when the user selects option "2" in the appointment menu to display their current appointments
// This function queries the database using the user's previously submitted ID to fetch and save in memory any appointments tied to that user.
// Any appointments in the database are returned in a list format.
func getAppointmentsByUserID(db *sql.DB, userID string) ([]appointment, error) {
	rows, err := db.Query(
		`SELECT
			id,
			pet_name,
			pet_species,
			pet_age,
			pet_weight,
			vaccinated,
			appointment_type,
			vet_name,
			appointment_time
		FROM appointments
		WHERE user_id = $1`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []appointment

	for rows.Next() {
		var a appointment
		var p pet

		err := rows.Scan(
			&a.id,
			&p.name,
			&p.species,
			&p.age,
			&p.weightKg,
			&p.vaccinated,
			&a.appointmentType,
			&a.vet,
			&a.dateTime,
		)
		if err != nil {
			return nil, err
		}

		a.pet = p
		appointments = append(appointments, a)
	}

	return appointments, nil
}

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
	// TODO: Change this to use supported method
	input = strings.Title(input)
	return input, nil
}

// appointmentMenu is a function displays a menu screen to the user with 3 options.
// The option that the user selects is normalised and then passed to main().
func appointmentMenu(scanner *bufio.Scanner) string {
	fmt.Println("1. Create new appointment")
	fmt.Println("2. View existing appointments")
	fmt.Println("3. Delete an appointment")
	fmt.Println("4. Exit")
	fmt.Print("> ")

	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// getNumberofPets is a function that prompts the user to enter the number of pets they wish to book an appointment for.
// The input is converted into an integer type and then validated.
// If the input is invalid, an error is returned.
func getNumberofPets(scanner *bufio.Scanner) (int, error) {
	var petCount int

	fmt.Println("Please enter how many pets you are booking appointments for: ")
	scanner.Scan()
	petCount, _ = strconv.Atoi(scanner.Text())

	if petCount > 0 && petCount <= 20 {
		return petCount, nil
	} else if petCount > 20 {
		return 0, fmt.Errorf("input exceeds limit of 20")
	} else {
		return 0, fmt.Errorf("input must have value between 0 and 20")
	}
}

// getSpecies is a helper function that prompts the user to provide their pet's species and lists available options using the "allowedSpecies" list.
// The input is stored and normalised.
// If the input is not listed in "allowedSpecies", the user is prompted again.
func getSpecies(scanner *bufio.Scanner) (string, error) {
	fmt.Println("Please enter pet species: ")

	for i, v := range allowedSpecies {
		fmt.Printf("%d. %s\n", i+1, v)
	}
	fmt.Print("> ")

	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(allowedSpecies) {
		return "", fmt.Errorf("please select one of the species displayed")
	}

	return allowedSpecies[choice-1], nil
}

// getAge is a helper function that prompts the user for their pet's age and stores it.
// The stored age is converted to an integer type.
// The name is passed through a validation check.
// If validation fails, an error is returned.
func getAge(scanner *bufio.Scanner) (int, error) {
	var input int

	fmt.Println("Please enter pet age: ")
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
func getWeightKg(scanner *bufio.Scanner) (float64, error) {
	var input float64

	fmt.Println("Please enter pet weight (Kg): ")
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
func getVaccinationStatus(scanner *bufio.Scanner) (bool, error) {
	var input string

	fmt.Println("Is pet vaccinated? (y/n): ")
	scanner.Scan()
	input = scanner.Text()

	switch input {
	case "y", "Y":
		return true, nil
	case "n", "N":
		return false, nil
	default:
		return false, fmt.Errorf("input must be y/n")
	}
}

// getAppointmentType is a helper function that prompts the user to choose an appointment type and lists available options using the "allowedAppointmentTypes" list.
// The input is stored and normalised.
// If the input is not listed in "allowedAppointmentTypes", the user is prompted again.
func getAppointmentType(scanner *bufio.Scanner) (string, error) {
	fmt.Println("Please enter appointment type for")

	for i, v := range allowedAppointmentTypes {
		fmt.Printf("%d. %s\n", i+1, v)
	}
	fmt.Print("> ")

	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(allowedAppointmentTypes) {
		return "", fmt.Errorf("please select one of the appointment types displayed")
	}

	return allowedAppointmentTypes[choice-1], nil
}

// getVet is a helper function that prompts the user to choose a preferred vet for their appointment and lists available options using the "allowedVets" list.
// The input is stored and normalised.
// If the input is not listed in "allowedAppointmentTypes", the user is prompted again.
func getVet(scanner *bufio.Scanner) (string, error) {
	fmt.Println("Please choose preferred vet for appointment")

	for i, v := range allowedVets {
		fmt.Printf("%d. %s\n", i+1, v)
	}
	fmt.Print("> ")

	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(allowedVets) {
		return "", fmt.Errorf("please select one of the vets displayed")
	}

	return allowedVets[choice-1], nil
}

// getPreferredDateTime is a helper function that allows the user to enter a preferred date and time for their appointment.
// The user is prompted for a date and time in a specified format.
// The input is stored and normalised.
// The input is parsed and converted into a date and time format.
// The input is then validated and an error is displayed if it doesn't pass the validation checks.
func getPreferredDateTime(scanner *bufio.Scanner) (time.Time, error) {
	fmt.Println("Please enter preferred date and time for appointment")
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

// getAppointment calls the helper functions repeatedly until a valid input is received from the user for all fields. This procedure is iterated for each appointment the user filled in details for.
// If an error is received for a helper function, getAppointment calls the function again, and the user is prompted for a valid input.
// If a valid input is received for a helper function, getAppointment will pass the valid input to the corresponding field in the newly initialised "appointment" objects.
// The appointment objects are stored in a list to accommodate multiple appointments.
// Once all fields in "appointment" are filled, getAppointment returns the list of "appointment" objects.
func getAppointment(scanner *bufio.Scanner) []appointment {
	appointments := make([]appointment, 0)

	var d pet

	for {
		name, err := getName(scanner, "Please enter pet name")
		if err == nil {
			d.name = name
			break
		}
		fmt.Println("Error:", err)
	}

	for {
		breed, err := getSpecies(scanner)
		if err == nil {
			d.species = breed
			break
		}
		fmt.Println("Error:", err)
	}

	for {
		age, err := getAge(scanner)
		if err == nil {
			d.age = age
			break
		}
		fmt.Println("Error:", err)
	}

	for {
		weightKg, err := getWeightKg(scanner)
		if err == nil {
			d.weightKg = weightKg
			break
		}
		fmt.Println("Error:", err)
	}

	for {
		vaccinated, err := getVaccinationStatus(scanner)
		if err == nil {
			d.vaccinated = vaccinated
			break
		}
		fmt.Println("Error:", err)
	}

	var a appointment

	for {
		appointmentType, err := getAppointmentType(scanner)
		if err == nil {
			a.appointmentType = appointmentType
			break
		}
		fmt.Println("Error:", err)
	}

	for {
		v, err := getVet(scanner)
		if err == nil {
			a.vet = v
			break
		}
		fmt.Println("Error:", err)
	}

	for {
		dt, err := getPreferredDateTime(scanner)
		if err == nil {
			a.dateTime = dt
			break
		}
		fmt.Println("Error:", err)
	}

	a.pet = d
	appointments = append(appointments, a)

	return appointments
}

// TODO: change to prevent user from deleting appointments that don't belong to them
func deleteAppointment(scanner *bufio.Scanner, db *sql.DB) error {
	fmt.Println("Enter appointment ID to delete: ")
	fmt.Print(">")

	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	id, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("Invalid ID: %v", err)
	}

	result, err := db.Exec("DELETE FROM appointments WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("Database error: %v:", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("No appointment found with ID %d", id)
	}

	fmt.Printf("Appointment %d deleted successfully\n", id)
	return nil
}

// summaryString prints a summary of each appointment's details.
func (a *appointment) summaryString(i int) string {
	var s string
	s = "-------------------------------------\n"
	s += fmt.Sprintf("Appointment %d information:\n", i)
	s += fmt.Sprintf("ID: %s\n", a.id)
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
