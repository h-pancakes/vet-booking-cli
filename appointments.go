package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Used for both user and pet name capture
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

func getAppointmentMenuChoice(scanner *bufio.Scanner) string {
	fmt.Println("1. Create new appointment")
	fmt.Println("2. View existing appointments")
	fmt.Println("3. Delete an appointment")
	fmt.Println("4. Exit")
	fmt.Print("> ")

	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

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

func gatherAppointmentInfo(scanner *bufio.Scanner) (
	petName, petSpecies string,
	petAge int,
	petWeightKg float64,
	vaccinated bool,
	appointmentType, vet string,
	appointmentTime time.Time,
) {

	var err error

	for {
		petName, err = getName(scanner, "Please enter pet name")
		if err == nil {
			break
		}
	}

	for {
		petSpecies, err = getSpecies(scanner)
		if err == nil {
			break
		}
	}

	for {
		petAge, err = getAge(scanner)
		if err == nil {
			break
		}
	}

	for {
		petWeightKg, err = getWeightKg(scanner)
		if err == nil {
			break
		}
	}

	for {
		vaccinated, err = getVaccinationStatus(scanner)
		if err == nil {
			break
		}
	}

	for {
		appointmentType, err = getAppointmentType(scanner)
		if err == nil {
			break
		}
	}

	for {
		vet, err = getVet(scanner)
		if err == nil {
			break
		}
	}

	for {
		appointmentTime, err = getPreferredDateTime(scanner)
		if err == nil {
			break
		}
	}

	return petName, petSpecies, petAge, petWeightKg, vaccinated, appointmentType, vet, appointmentTime
}

// This function exists to hold unchangeable business logic which inevitably includes already-used CLI validation logic. It makes an appointment OBJECT.
func NewAppointment(appointmentType, vet, petName, petSpecies string, petAge int, petWeightKg float64, petVaccinated bool, dateTime time.Time) (appointment, error) {

	// TODO: primitive appointment clashing logic here

	nameLength := 0
	for _, character := range petName {
		if character != ' ' {
			nameLength++
		}
	}
	if nameLength < 1 {
		return appointment{}, fmt.Errorf("name must be at least 1 character")
	}
	if nameLength > 20 {
		return appointment{}, fmt.Errorf("character limit is 20 characters")
	}
	for _, character := range petName {
		if character >= 'A' && character <= 'Z' {
			continue
		}
		if character >= 'a' && character <= 'z' {
			continue
		}
		if character == ' ' {
			continue
		}
		if character == '-' {
			continue
		}
		return appointment{}, fmt.Errorf("name can only contain A-Z, hyphens, and spaces")
	}

	validSpecies := false
	for _, allowed := range allowedSpecies {
		if petSpecies == allowed {
			validSpecies = true
			break
		}
	}
	if !validSpecies {
		return appointment{}, fmt.Errorf("please select one of the species displayed")
	}

	if petAge < 0 || petAge > 30 {
		return appointment{}, fmt.Errorf("age must be between 0 and 30 years")
	}
	if petWeightKg < 1 || petWeightKg > 120 {
		return appointment{}, fmt.Errorf("weight must be between 1 and 120kg")
	}

	validAppointmentType := false
	for _, allowed := range allowedAppointmentTypes {
		if appointmentType == allowed {
			validAppointmentType = true
			break
		}
	}
	if !validAppointmentType {
		return appointment{}, fmt.Errorf("please select one of the appointment types displayed")
	}

	validVet := false
	for _, allowed := range allowedVets {
		if vet == allowed {
			validVet = true
			break
		}
	}
	if !validVet {
		return appointment{}, fmt.Errorf("please select one of the vets displayed")
	}

	if dateTime.Before(time.Now()) {
		return appointment{}, fmt.Errorf("Appointment cannot be in the past")
	}

	return appointment{appointmentType: appointmentType, vet: vet, petName: petName, petSpecies: petSpecies, petAge: petAge, petWeightKg: petWeightKg, petVaccinated: petVaccinated, dateTime: dateTime}, nil
}

// func deleteAppointment(scanner *bufio.Scanner, db *sql.DB) error {
// 	fmt.Println("Enter appointment ID to delete: ")
// 	fmt.Print(">")

// 	scanner.Scan()
// 	input := strings.TrimSpace(scanner.Text())

// 	id, err := strconv.Atoi(input)
// 	if err != nil {
// 		return fmt.Errorf("Invalid ID: %v", err)
// 	}

// 	result, err := db.Exec("DELETE FROM appointments WHERE id = $1", id)
// 	if err != nil {
// 		return fmt.Errorf("Database error: %v:", err)
// 	}

// 	rowsAffected, _ := result.RowsAffected()
// 	if rowsAffected == 0 {
// 		return fmt.Errorf("No appointment found with ID %d", id)
// 	}

// 	fmt.Printf("Appointment %d deleted successfully\n", id)
// 	return nil
// }

// summaryString prints a summary of each appointment's details.
func (a *appointment) summaryString(i int) string {
	var s string
	s = "-------------------------------------\n"
	s += fmt.Sprintf("Appointment %d information:\n", i)
	s += fmt.Sprintf("Pet Name: %s\n", a.petName)
	s += fmt.Sprintf("Species: %s\n", a.petSpecies)
	s += fmt.Sprintf("Age: %d\n", a.petAge)
	s += fmt.Sprintf("Weight (kg): %.2f\n", a.petWeightKg)
	s += fmt.Sprintf("Vaccinated?: %t\n", a.petVaccinated)
	s += fmt.Sprintf("Appointment Type: %s\n", a.appointmentType)
	s += fmt.Sprintf("Vet: %s\n", a.vet)
	s += fmt.Sprintf("Appointment Date & Time: %s\n", a.dateTime.Format("Monday, 02 Jan 2006 at 15:04"))
	s += "-------------------------------------\n"

	return s
}

func promptUserDeleteAppointment(scanner *bufio.Scanner) (int, error) {
	fmt.Println("Enter appointment number to delete:")
	fmt.Print(">")

	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	appointmentNumber, err := strconv.Atoi(input)
	if err != nil {
		return -67, ErrInvalidAppointmentNumber
	}

	userChoice := appointmentNumber - 1

	return userChoice, nil
}
