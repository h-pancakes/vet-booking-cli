package main

import "database/sql"

type AppointmentsRepository interface {
	CreateAppointment(userID string, a appointment) error
	RetrieveAppointments(userID string) ([]appointment, error)
	UpdateAppointment(id, userID string, a appointment) error
	DeleteAppointment(id, userID string) error
}

type AppointmentRepo struct {
	DB *sql.DB
}

func NewAppointmentRepo(db *sql.DB) *AppointmentRepo {
	return &AppointmentRepo{DB: db}
}

func (r *AppointmentRepo) CreateAppointment(userID string, a appointment) error {

	query := "INSERT INTO appointments (user_id, pet_name, pet_species, pet_age, pet_weight, vaccinated, appointment_type, vet_name, appointment_time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	_, err := r.DB.Exec(query, userID, a.petName, a.petSpecies, a.petAge, a.petWeightKg, a.petVaccinated, a.appointmentType, a.vet, a.dateTime)

	if err != nil {
		return ErrDatabaseQueryFailure
	}

	return nil
}

func (r *AppointmentRepo) RetrieveAppointments(userID string) ([]appointment, error) {

	query := "SELECT id, user_id, pet_name, pet_species, pet_age, pet_weight, vaccinated, appointment_type, vet_name, appointment_time FROM appointments WHERE user_id = $1"

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, ErrDatabaseQueryFailure
	}
	defer rows.Close()

	var appointments []appointment
	for rows.Next() {
		var a appointment
		err = rows.Scan(&a.id, &a.userID, &a.petName, &a.petSpecies, &a.petAge, &a.petWeightKg, &a.petVaccinated, &a.appointmentType, &a.vet, &a.dateTime)
		if err != nil {
			return nil, ErrDatabaseScanFailure
		}
		appointments = append(appointments, a)
	}
	err = rows.Err()
	if err != nil {
		return nil, ErrDatabaseRowIterationFailure
	}

	return appointments, nil
}

func (r *AppointmentRepo) UpdateAppointment(id, userID string, a appointment) error {

	query := "UPDATE appointments SET pet_name = $1, pet_species = $2, pet_age = $3, pet_weight = $4, vaccinated = $5, appointment_type = $6, vet_name = $7, appointment_time = $8 WHERE id = $9 AND user_id = $10"

	_, err := r.DB.Exec(query, a.petName, a.petSpecies, a.petAge, a.petWeightKg, a.petVaccinated, a.appointmentType, a.vet, a.dateTime, a.id, a.userID)

	if err != nil {
		return ErrDatabaseQueryFailure
	}

	return nil
}

func (r *AppointmentRepo) DeleteAppointment(id, userID string) error {

	query := "DELETE FROM appointments WHERE id = $1 AND user_id = $2"

	//  TODO: complete this error handling
	_, err := r.DB.Exec(query, id, userID)
	if err != nil {
		return err
	}

	return nil
}
