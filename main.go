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

	for i := 0; i < dogCount; i++ {
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

		fmt.Println("Enter dog", i+1, "breed: ")
		scanner.Scan()
		dogs[i].breed = scanner.Text()

		fmt.Println("Is dog", i+1, "vaccinated? (y/n): ")
		scanner.Scan()
		input = scanner.Text()

		switch input {
		case "y":
			dogs[i].vaccinated = true
		case "n":
			dogs[i].vaccinated = false
		default:
			fmt.Println("Invalid input!")
			return nil
		}
	}

	return dogs
}

// appointmentDecision returns a message based on vaccination status.
func (d *dog) appointmentDecision() string {
	if d.vaccinated {
		return fmt.Sprintf(
			"%s (%s) is vaccinated! Booking regular checkup appointment...", d.name, d.breed,
		)
	}
	return fmt.Sprintf(
		"%s (%s) is NOT vaccinated! Please book a vaccination appointment first", d.name, d.breed,
	)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	dogCount := dogCounter(scanner)
	dogs := bookAppointment(scanner, dogCount)

	for _, d := range dogs {
		fmt.Println(d.appointmentDecision())
	}
}
