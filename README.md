# vet-booking-cli

A menu-driven CLI program written in Go.
Users can log in via email and password, create, read, and delete veterinary appointments for their pets.

# Requirements

- Go 1.20+
- PostgreSQL

# Setup

1. Create a PostgreSQL database:
CREATE DATABASE vet_booking;

2. Connect to the database:
\c vet_booking

3. Run the schema:
 - psql vet_booking < schema.sql (Powershell)
 - \i schema.sql (Bash - make sure current directory is project root directory beforehand)

4. Create a .env file in the project root directory using .env.example as a template:
- cp .env.example .env (Linux/MacOS)
- Copy-Item .env.example .env (Windows)

5. Edit .env and replace username and password with YOUR PostgreSQL credentials:
DATABASE_URL=postgres://username:password@localhost:5432/vet_booking?sslmode=disable

6. Change directory to project root

7. Run program:
go run .

# Known limitations

 - Business logic is not fully implemented yet e.g. appointment clashing
 - Error handling is inconsistent/missing in certain parts like legacy code parts (scanner errors)
 - Complete refactoring to new layered architecture currently only limited to user service, not appointment service
 - Tests are present but are not extensive enough and dont cover crucial parts of the application
 - Major bug in DELETE function that doesn't prevent a user from deleting any appointment, regardless of whether they own it or not i.e. using an older placeholder method of deletion, not to mention mixing db logic and business logic
 - Current layout of code across files is messy and lacks splitting into directories and packages
 - Some CLI functions like getUserPhone and getUserEmail are not uniform with similar get functions as they also contain the prompt instead of the prompt being passed in as a parameter
 - Failed appointment save in createNewAppointment (legacy function) exits the program on failed database save
 - Email validation is not extensive enough (perhaps will replace with regex)


 - Program is not containerised yet and so is tedious to install and run
