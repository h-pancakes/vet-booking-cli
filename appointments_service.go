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
