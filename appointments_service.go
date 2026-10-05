package main

type AppointmentService struct {
	repo AppointmentsRepository
}

func NewAppointmentService(r AppointmentsRepository) *AppointmentService {
	return &AppointmentService{repo: r}
}

func (s *AppointmentService) CreateAppointment()
