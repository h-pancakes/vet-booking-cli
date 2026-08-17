package main

import "time"

// user is a struct that holds information about the user of the booking service.
// It contains contact details persisted to the database.
type user struct {
	firstName string
	lastName  string
	phone     string
	email     string
}

// pet is a struct that holds information about a pet that the user is booking an appointment for.
// This information is stored in the appointments table in the database.
type pet struct {
	name       string
	species    string
	age        int
	weightKg   float64
	vaccinated bool
}

// appointment is a struct that holds all information related to an appointment booked by the user.
// This information is stored in the appointments table in the database.
type appointment struct {
	id              string
	appointmentType string
	pet             pet
	vet             string
	dateTime        time.Time
}
