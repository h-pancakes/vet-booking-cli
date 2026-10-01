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
	return &fakeRepo{byEmail: map[string]user{}}
}

func (f *fakeRepo) Create(u user) (user, error) {
	f.nextId++
	u.id = strconv.Itoa(f.nextId)
	f.byEmail[u.phone] = u
	return u, nil
}

func (f *fakeRepo) FindByEmail(phone string) (user, error) {
	u, ok := f.byEmail[phone]
	if !ok {
		return user{}, ErrEmailNotFound
	}
	return u, nil
}

func TestRegisterUser_Success(t *testing.T) {
	svc := NewUserService(newFakeRepo())
	got, err := svc.RegisterUser("Bob", "Builder", "07668547589", "bobbuilder@hotmail.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.id == "0" {
		t.Error("expected id to be assigned")
	}
}
