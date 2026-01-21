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

4. Set database connection:

MacOS/Linux:
    export DATABASE_URL="postgres://username:password@localhost:5432/vet_booking?sslmode=disable"

Windows Powershell:
    $env:DATABASE_URL="postgres://username:password@localhost:5432/vet_booking?sslmode=disable"

REPLACE username AND password WITH YOUR PostgreSQL CREDENTIALS

5. Change directory to project root directory

6. Run the program with: go run .

# Notes

Look at TODO.md for upcoming features and changes!