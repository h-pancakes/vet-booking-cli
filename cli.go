package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func runCLI(UserService *UserService, AppointmentService *AppointmentService) bool {

	scanner := bufio.NewScanner(os.Stdin)
	var CurrentUser user
	// not ideal?:
	var err error

first:
	for {
		choice := getMainMenuChoice(scanner)

		switch choice {
		case "1":
			for {
				var password []byte
				firstName, lastName, phone, email, password, err := gatherUserInfo(scanner)

				if err != nil {
					fmt.Println("Error", err)
				}

				_, err = UserService.RegisterUser(
					firstName, lastName, phone, email, password,
				)
				if err != nil {
					fmt.Println("Error:", err)
					continue
				}
				fmt.Println("Account created successfully!")
				break
			}

		case "2":
			for {
				email, password, err := promptUserLogin(scanner)

				if err != nil {
					fmt.Println("Error: ", err)
					continue
				}

				LoggedInUser, err := UserService.LoginUser(email, password)
				if err != nil {
					if errors.Is(err, ErrInvalidEmailOrPassword) {
						fmt.Println(ErrInvalidEmailOrPassword)
					} else {
						fmt.Println("Login failed:", err)
					}
					continue
				}
				// Why is there currentUser AND loggedInUser???
				CurrentUser = LoggedInUser
				fmt.Println("Welcome,", CurrentUser.firstName)
				break first
			}

		case "3":
			fmt.Println("Goodbye!")
			return true

		default:
			fmt.Println("Invalid option, please try again.")
			continue
		}

	}

	for {
		userChoice := getAppointmentMenuChoice(scanner)

		switch userChoice {

		// Option to create new appointment
		case "1":

			petName, petSpecies, petAge, petWeightKg, petVaccinated, appointmentType, vet, appointmentTime := gatherAppointmentInfo(scanner)

			// TODO: Add error handling to gatherAppointmentInfo

			err = AppointmentService.BookAppointment(CurrentUser.id, appointmentType, vet, petName, petSpecies, petAge, petWeightKg, petVaccinated, appointmentTime)

			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Println("Appointment booked successfully!")

		// Option to view apppointments
		case "2":
			appointments, err := AppointmentService.ViewAppointments(CurrentUser.id)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Println(CurrentUser.UserSummaryString())

			for i, a := range appointments {
				fmt.Println(a.summaryString(i + 1))
			}

		// Option to update an appointment
		case "3":
			appointments, err := AppointmentService.ViewAppointments(CurrentUser.id)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			for i, a := range appointments {
				fmt.Println(a.summaryString(i + 1))
			}

			index, err := promptUserUpdateAppointment(scanner)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if index < 0 || index >= len(appointments) {
				fmt.Println("Error:", ErrInvalidAppointmentNumber)
				continue
			}

			chosenAppointment := appointments[index]

			petName, petSpecies, petAge, petWeightKg, petVaccinated, appointmentType, vet, appointmentTime := gatherAppointmentInfo(scanner)

			err = AppointmentService.UpdateAppointment(chosenAppointment.id, CurrentUser.id, appointmentType, vet, petName, petSpecies, petAge, petWeightKg, petVaccinated, appointmentTime)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Println("Appointment updated successfully!")

		// Option to delete an appointment
		case "4":
			appointments, err := AppointmentService.ViewAppointments(CurrentUser.id)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			for i, a := range appointments {
				fmt.Println(a.summaryString(i + 1))
			}

			index, err := promptUserDeleteAppointment(scanner)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if index < 0 || index >= len(appointments) {
				fmt.Println("Error:", ErrInvalidAppointmentNumber)
				continue
			}

			chosenAppointment := appointments[index]

			err = AppointmentService.RemoveAppointment(chosenAppointment.id, CurrentUser.id)

			if err != nil {
				fmt.Println("Error:", err)
			}

			fmt.Println("Appointment deleted successfully!")

		// Option to exit menu and end program
		case "5":
			fmt.Println("Goodbye!")
			return true

		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
