package main

import (
	"strconv"
	"testing"
)

type fakeRepo struct {
	byEmail map[string]user
	nextId  int
}

func newFakeRepo() *fakeRepo {
	return &fakeRepo{
		byEmail: map[string]user{},
	}
}

func (f *fakeRepo) CreateUser(u user) error {
	f.nextId++
	u.id = strconv.Itoa(f.nextId)
	f.byEmail[u.email] = u
	return nil
}

func (f *fakeRepo) FindByEmail(email string) (user, error) {
	u, ok := f.byEmail[email]
	if !ok {
		return user{}, ErrEmailNotFound
	}
	return u, nil
}

func (f *fakeRepo) FindByEmailForLogin(email string) (user, []byte, error) {
	u, ok := f.byEmail[email]
	if !ok {
		return user{}, nil, ErrEmailNotFound
	}
	return u, []byte(u.passwordHash), nil
}

func TestRegisterUser_Success(t *testing.T) {
	svc := NewUserService(newFakeRepo())

	password := []byte("password123")
	got, err := svc.RegisterUser(
		"Bob",
		"Builder",
		"07668547589",
		"bobbuilder@hotmail.com",
		password,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.id == "0" {
		t.Error("expected id to be assigned")
	}
}
