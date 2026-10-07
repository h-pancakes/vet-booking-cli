package main

import (
	"fmt"
	"time"
)

type AppointmentService struct {
	repo AppointmentsRepository
}

func NewAppointmentService(r AppointmentsRepository) *AppointmentService {
	return &AppointmentService{repo: r}
}

func (s *AppointmentService) BookAppointment(userID, appointmentType, vet, petName, petSpecies string, petAge int, petWeightKg float64, petVaccinated bool, appointmentTime time.Time) (appointment, error) {

	a, err := NewAppointment(appointmentType, vet, petName, petSpecies, petAge, petWeightKg, petVaccinated, appointmentTime)

	if err != nil {
		return appointment{}, err
	}

	err = s.repo.CreateAppointment(userID, a)
	if err != nil {
		return appointment{}, fmt.Errorf("database save failed: %w", err)
	}

	return a, nil
}

func (s *AppointmentService) ViewAppointments(userID string) ([]appointment, error) {
	appointments, err := s.repo.RetrieveAppointments(userID)
	if err != nil {
		return nil, err
	}

	if len(appointments) == 0 {
		return nil, ErrNoAppointmentsFound
	}

	return appointments, err
}

func (s AppointmentService) RemoveAppointment(appointment appointment) error {

	err := s.repo.DeleteAppointment(appointment.id, appointment.userID)

	if err != nil {
		return err
	}

	return nil
}

// This function exists to hold unchangeable business logic which inevitably includes already-used CLI validation logic. It makes an appointment OBJECT.
func NewAppointment(appointmentType, vet, petName, petSpecies string, petAge int, petWeightKg float64, petVaccinated bool, dateTime time.Time) (appointment, error) {

	// TODO: primitive appointment clashing logic here

	nameLength := 0
	for _, character := range petName {
		if character != ' ' {
			nameLength++
		}
	}
	if nameLength < 1 {
		return appointment{}, fmt.Errorf("name must be at least 1 character")
	}
	if nameLength > 20 {
		return appointment{}, fmt.Errorf("character limit is 20 characters")
	}
	for _, character := range petName {
		if character >= 'A' && character <= 'Z' {
			continue
		}
		if character >= 'a' && character <= 'z' {
			continue
		}
		if character == ' ' {
			continue
		}
		if character == '-' {
			continue
		}
		return appointment{}, fmt.Errorf("name can only contain A-Z, hyphens, and spaces")
	}

	validSpecies := false
	for _, allowed := range allowedSpecies {
		if petSpecies == allowed {
			validSpecies = true
			break
		}
	}
	if !validSpecies {
		return appointment{}, fmt.Errorf("please select one of the species displayed")
	}

	if petAge < 0 || petAge > 30 {
		return appointment{}, fmt.Errorf("age must be between 0 and 30 years")
	}
	if petWeightKg < 1 || petWeightKg > 120 {
		return appointment{}, fmt.Errorf("weight must be between 1 and 120kg")
	}

	validAppointmentType := false
	for _, allowed := range allowedAppointmentTypes {
		if appointmentType == allowed {
			validAppointmentType = true
			break
		}
	}
	if !validAppointmentType {
		return appointment{}, fmt.Errorf("please select one of the appointment types displayed")
	}

	validVet := false
	for _, allowed := range allowedVets {
		if vet == allowed {
			validVet = true
			break
		}
	}
	if !validVet {
		return appointment{}, fmt.Errorf("please select one of the vets displayed")
	}

	if dateTime.Before(time.Now()) {
		return appointment{}, fmt.Errorf("Appointment cannot be in the past")
	}

	return appointment{appointmentType: appointmentType, vet: vet, petName: petName, petSpecies: petSpecies, petAge: petAge, petWeightKg: petWeightKg, petVaccinated: petVaccinated, dateTime: dateTime}, nil
}
