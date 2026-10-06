package main

import "time"

// holds information about the pet owner
type user struct {
	id           string
	firstName    string
	lastName     string
	phone        string
	email        string
	passwordHash string
}

// holds appointment information
type appointment struct {
	id              string
	userID          string
	appointmentType string
	vet             string
	dateTime        time.Time
	petName         string
	petSpecies      string
	petAge          int
	petWeightKg     float64
	petVaccinated   bool
}
