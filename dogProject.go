package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type Dog struct {
	name string
	breed string
	vaccinated bool
}

func (d *Dog) appointmentDecision() string {
	if d.vaccinated {
		return fmt.Sprintf("Thanks for booking your %s, %s for a checkup appointment! You will receive an email with the details. We look forward to seeing you and %s soon!", d.breed, d.name, d.name)
	} else {
		return fmt.Sprintf("Sorry! We cannot book an appointment without having %s vaccinated. Please book a vaccination appointment.", d.name)
	}
}

// I was working on this function
func bookAppointment(scanner *bufio.Scanner) {
	dogCount := dogCounter()
	dogs := make([]Dog, dogCount)
	var input string

	for i := 0; i < dogCount; i++ {
		fmt.Println("Enter dog", i+1, "name: ")
		scanner.Scan()
		dogs[i].name = scanner.Text()

		fmt.Println("Enter dog", i+1, "breed: ")
		scanner.Scan()
		dogs[i].breed = scanner.Text()
		
		fmt.Println("Is dog", i+1, "vaccinated? (y/n): ")
		scanner.Scan()
		input = scanner.Text()

		if input == "y" {
			dogs[i].vaccinated = true
		} else if input == "n" {
			dogs[i].vaccinated = false
		} else {
			fmt.Println("Invalid input!")
			return
		}	
	}

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
	bookAppointment(scanner)
}
