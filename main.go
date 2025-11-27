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

// dogCounter prompts the user to enter the number of dogs they wish to book an appointment for.
func dogCounter(scanner *bufio.Scanner) int {
	var dogCount int

	fmt.Println("Welcome to our booking service!")
	fmt.Println("Please enter how many dogs you are booking in today: ")
	scanner.Scan()
	dogCount, _ = strconv.Atoi(scanner.Text())

	if dogCount > 0 && dogCount <= 20 {
		return dogCount
	} else {
		panic("invalid input")
	}
}

// bookAppointment collects user input to create a slice of Dogs.
func bookAppointment(scanner *bufio.Scanner, dogCount int) []dog {
	dogs := make([]dog, dogCount)

	for i := 0; i < dogCount; i++ { // Need to polish validation here
		var input string

		fmt.Println("Enter dog", i+1, "name: ")
		scanner.Scan()

		input = scanner.Text()

		input = strings.TrimSpace(input)

		if len(input) < 1 {
			panic("name too short")
		}

		if len(input) > 20 {
			panic("name too long")
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
			panic("invalid input")
		}

		dogs[i].name = input

		fmt.Println("Enter dog", i+1, "breed: ") // Need to add validation here
		scanner.Scan()
		dogs[i].breed = scanner.Text()

		fmt.Println("Enter dog", i+1, "age: ") // Need to add validation here
		scanner.Scan()
		dogs[i].age, _ = strconv.Atoi(scanner.Text())

		fmt.Println("Enter dog", i+1, "weight: ") // Need to add validation here
		scanner.Scan()
		dogs[i].weightKg, _ = strconv.ParseFloat(scanner.Text(), 64)

		fmt.Println("Is dog", i+1, "vaccinated? (y/n): ")
		scanner.Scan()
		input = scanner.Text()

		switch input {
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
