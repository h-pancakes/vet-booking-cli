package main

import "database/sql"

type AppointmentsRepository interface {
}

type AppointmentRepo struct {
	DB *sql.DB
}

func NewAppointmentRepo(db *sql.DB) *AppointmentRepo {
	return &AppointmentRepo{DB: db}
}
