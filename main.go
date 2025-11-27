package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// dog holds information about a pet being booked for an appointment.
type dog struct {
	name       string
	breed      string
	age        int
	weightKg   float64
	vaccinated bool
}

var allowedBreeds = []string{
	"Labrador",
	"Poodle",
	"Beagle",
	"Bulldog",
	"Dachshund",
}

// dogCounter prompts the user to enter the number of dogs they wish to book an appointment for.
func dogCounter(scanner *bufio.Scanner) int {
	var dogCount int

	fmt.Println("Welcome to our booking service!")
	fmt.Println("Please enter how many dogs you are booking in today: ")
	scanner.Scan()
	dogCount, _ = strconv.Atoi(scanner.Text())

	if dogCount > 0 && dogCount <= 20 { // Need to add error handling here
		return dogCount
	} else {
		panic("invalid amount")
	}
}

// bookAppointment collects user input to create a slice of Dogs.
func bookAppointment(scanner *bufio.Scanner, dogCount int) []dog {
	dogs := make([]dog, dogCount)

	for i := 0; i < dogCount; i++ { /* Need to add error handling here */
		var strInput string
		var intInput int
		var floatInput float64

		fmt.Println("Enter dog", i+1, "name: ")
		scanner.Scan()

		strInput = scanner.Text()

		strInput = strings.TrimSpace(strInput)

		trimmed := strings.ReplaceAll(strInput, " ", "")

		if len(trimmed) < 1 {
			panic("name too short")
		}

		if len(trimmed) > 20 {
			panic("name too long")
		}

		for _, c := range strInput {
			if c >= 'A' && c <= 'Z' {
				continue
			}
			if c >= 'a' && c <= 'z' {
				continue
			}
			if c == ' ' {
				continue
			}
			panic("invalid name")
		}

		dogs[i].name = strInput

		fmt.Println("Enter dog", i+1, "breed: ") /* Need to add error handling here */
		fmt.Println("Allowed Breeds:", allowedBreeds)
		scanner.Scan()
		strInput = scanner.Text()
		strInput = strings.TrimSpace(strInput)
		strInput = strings.Title(strInput)
		valid := false
		for _, breed := range allowedBreeds {
			if breed == strInput {
				valid = true
				break
			}
		}

		if !valid {
			panic("invalid breed")
		}

		dogs[i].breed = strInput

		fmt.Println("Enter dog", i+1, "age: ") /* Need to add error handling here */
		scanner.Scan()
		intInput, _ = strconv.Atoi(scanner.Text())
		if intInput < 0 || intInput > 30 {
			panic("invalid age")
		}

		dogs[i].age = intInput

		fmt.Println("Enter dog", i+1, "weight: ") /* Need to add error handling here */
		scanner.Scan()
		floatInput, _ = strconv.ParseFloat(scanner.Text(), 64)
		if floatInput < 1 || floatInput > 120 {
			panic("invalid weight")
		}

		dogs[i].weightKg = floatInput

		fmt.Println("Is dog", i+1, "vaccinated? (y/n): ")
		scanner.Scan()
		strInput = scanner.Text()

		switch strInput {
		case "y":
			dogs[i].vaccinated = true
		case "n":
			dogs[i].vaccinated = false
		case "Y":
			dogs[i].vaccinated = true
		case "N":
			dogs[i].vaccinated = false
		default:
			fmt.Println("Invalid input!")
			return nil
		}
	}

	return dogs
}

// summaryString prints a summary of each dog's details
func (d *dog) summaryString(i int) string {
	var s string
	s = "-------------------------------------\n"
	s += fmt.Sprintf("Dog %d summary:\n", i)
	s += fmt.Sprintf("Name: %s\n", d.name)
	s += fmt.Sprintf("Breed: %s\n", d.breed)
	s += fmt.Sprintf("Age: %d\n", d.age)
	s += fmt.Sprintf("Weight (kg): %.2f\n", d.weightKg)
	s += fmt.Sprintf("Vaccinated?: %t\n", d.vaccinated)
	s += "-------------------------------------\n"

	return s
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	dogCount := dogCounter(scanner)
	dogs := bookAppointment(scanner, dogCount)

	for i, d := range dogs {
		fmt.Println(d.summaryString(i + 1))
	}
}
