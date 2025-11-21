package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Dog struct {
	name       string
	breed      string
	vaccinated bool
}

func (d *Dog) appointmentDecision() string {
	if d.vaccinated {
		return fmt.Sprintf(
			"%s (%s) is vaccinated! Booking regular checkup appointment...", d.name, d.breed,
		)
	}
	return fmt.Sprintf(
		"%s (%s) is NOT vaccinated! Please book a vaccination appointment first", d.name, d.breed,
	)
}

func bookAppointment(scanner *bufio.Scanner) []Dog {
	dogCount := dogCounter()
	dogs := make([]Dog, dogCount)

	for i := 0; i < dogCount; i++ {
		fmt.Println("Enter dog", i+1, "name: ")
		scanner.Scan()
		dogs[i].name = scanner.Text()

		fmt.Println("Enter dog", i+1, "breed: ")
		scanner.Scan()
		dogs[i].breed = scanner.Text()

		fmt.Println("Is dog", i+1, "vaccinated? (y/n): ")
		scanner.Scan()
		input := scanner.Text()

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

func dogCounter(scanner *bufio.Scanner) int {
	var dogCount int

	fmt.Println("Hi and welcome to our booking service!")
	fmt.Println("Please enter how many dogs you are booking in today: ")
	scanner.Scan()
	dogCount, _ = strconv.Atoi(scanner.Text())

	return dogCount
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	dogcount := dogCounter(scanner)
	dogs := bookAppointment(scanner, dogcount)
}
