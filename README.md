# vet-booking-cli

A command line simple veterinary appointment booking system written in Go.
Users can log in, create and view appointments stored in a PostgreSQL database.

# Requirements

- Go 1.20+
- PostgreSQL

# Setup

1. Create a PostgreSQL database e.g. vet_booking

2. Connect to the database: \c vet_booking

3. Run the schema: psql vet_booking < schema.sql

4. Create a .env file in the project root directory using .env.example as a template:
- cp .env.example .env (Linux/MacOS)
- Copy-Item .env.example .env (Windows)

5. Edit .env and replace username and password with YOUR PostgreSQL credentials:
DATABASE_URL=postgres://username:password@localhost:5432/vet_booking?sslmode=disable
 

6. Change directory to project root

7. Run with:
go run .

# Notes

Check out TODO.md for upcoming features!