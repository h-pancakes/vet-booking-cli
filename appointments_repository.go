package main

import "database/sql"

type AppointmentsRepository interface {
	CreateAppointmentForUser(userID string, a appointment) error
	RetrieveAppointmentsForUser(userID string) ([]appointment, error)
	UpdateAppointmentForUser(id, userID string, a appointment) error
	DeleteAppointmentForUser(id, userID string) error
}

type AppointmentRepo struct {
	DB *sql.DB
}

func NewAppointmentRepo(db *sql.DB) *AppointmentRepo {
	return &AppointmentRepo{DB: db}
}

func (r *AppointmentRepo) CreateAppointmentForUser(userID string, a appointment) error {

	query := "INSERT INTO appointments (user_id, pet_name, pet_species, pet_age, pet_weight, vaccinated, appointment_type, vet_name, appointment_time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	_, err := r.DB.Exec(query, userID, a.petName, a.petSpecies, a.petAge, a.petWeightKg, a.petVaccinated, a.appointmentType, a.vet, a.dateTime)

	if err != nil {
		return ErrDatabaseQueryFailure
	}

	return nil
}

func (r *AppointmentRepo) RetrieveAppointmentsForUser(userID string) ([]appointment, error) {

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

func (r *AppointmentRepo) UpdateAppointmentForUser(id, userID string, a appointment) error {

	query := "UPDATE appointments SET pet_name = $1, pet_species = $2, pet_age = $3, pet_weight = $4, vaccinated = $5, appointment_type = $6, vet_name = $7, appointment_time = $8 WHERE id = $9 AND user_id = $10"

	result, err := r.DB.Exec(query, a.petName, a.petSpecies, a.petAge, a.petWeightKg, a.petVaccinated, a.appointmentType, a.vet, a.dateTime, id, userID)
	if err != nil {
		return ErrDatabaseQueryFailure
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return ErrDatabaseQueryFailure
	}
	if rowsAffected == 0 {
		return ErrUnauthorised
	}

	return nil
}

func (r *AppointmentRepo) DeleteAppointmentForUser(id, userID string) error {

	query := "DELETE FROM appointments WHERE id = $1 AND user_id = $2"

	result, err := r.DB.Exec(query, id, userID)
	if err != nil {
		return ErrDatabaseQueryFailure
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return ErrDatabaseQueryFailure
	}
	if rowsAffected == 0 {
		return ErrUnauthorised
	}

	return nil
}
