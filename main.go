package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// getMainMenuChoice is a function that displays a menu screen to the user with 3 options.
// The option that the user selects is normalised and then passed to main().
func getMainMenuChoice(scanner *bufio.Scanner) string {
	fmt.Println("1. New user")
	fmt.Println("2. Existing user")
	fmt.Println("3. Exit")
	fmt.Print("> ")

	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func main() {
	var currentUser *user

	scanner := bufio.NewScanner(os.Stdin)

	_ = godotenv.Load()
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		fmt.Println("DATABASE_URL environment variable not set")
		return
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	UserService := NewUserService(NewUserRepo(db))

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

				loggedInUser, err := UserService.LoginUser(email, password)
				if err != nil {
					if errors.Is(err, ErrInvalidEmailOrPassword) {
						fmt.Println(ErrInvalidEmailOrPassword)
					} else {
						fmt.Println("Login failed:", err)
					}
					continue
				}
				// This is not ideal but mix of pass by val and pass by pointer - Fix!
				currentUser = &loggedInUser
				fmt.Println("Welcome,", currentUser.firstName)
				break first
			}

		case "3":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option, please try again.")
			continue
		}

	}

	for {
		userChoice := getAppointmentMenuChoice(scanner)

		switch userChoice {
		case "1":

			appointmentInfo := gatherAppointmentInfo(scanner)

			isAppointmentCreated := createNewAppointment(db, currentUser, appointmentInfo)
			if isAppointmentCreated {
				return
			}

		case "2":
			appts, err := getAppointmentsByUserID(db, currentUser.id)
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}

			if len(appts) == 0 {
				fmt.Println("No appointments yet.")
				continue
			}

			fmt.Println(currentUser)
			for i, a := range appts {
				fmt.Println(a.summaryString(i + 1))
			}

		case "3":
			appts, err := getAppointmentsByUserID(db, currentUser.id)
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}

			if len(appts) == 0 {
				fmt.Println("No appointments to delete.")
				continue
			}

			for i, a := range appts {
				fmt.Println(a.summaryString(i + 1))
			}

			if err := deleteAppointment(scanner, db); err != nil {
				fmt.Println("Error:", err)
			}

		case "4":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
