package main

import "testing"

func TestNewUserFirstNameEmpty(t *testing.T) {
	_, err := NewUser("", "Smith", "07875166726", "johnsmith@gmail.com")
	if err != ErrNameTooShort {
		t.Errorf("Expected error: %s", ErrNameTooShort.Error())
	}
}

func TestNewUserLastNameEmpty(t *testing.T) {
	_, err := NewUser("John", "", "07875166726", "johnsmith@gmail.com")
	if err != ErrNameTooShort {
		t.Errorf("Expected error: %s", ErrNameTooShort.Error())
	}
}

func TestNewUserFirstNameTooLong(t *testing.T) {
	_, err := NewUser("Johnnyyyyyyyyyyyyyyyy", "Smith", "07875166726", "johnsmith@gmail.com")
	if err != ErrNameTooLong {
		t.Errorf("Expected error: %s", ErrNameTooLong.Error())
	}
}

func TestNewUserLastNameTooLong(t *testing.T) {
	_, err := NewUser("John", "Smithyyyyyyyyyyyyyyyy", "07875166726", "johnsmith@gmail.com")
	if err != ErrNameTooLong {
		t.Errorf("Expected error: %s", ErrNameTooLong.Error())
	}
}

func TestNewUserFirstNameHasInvalidCharacters(t *testing.T) {
	_, err := NewUser("J0hn", "Smith", "07875166726", "johnsmith@gmail.com")
	if err != ErrNameContainsInvalidCharacters {
		t.Errorf("Expected error: %s", ErrNameTooLong.Error())
	}
}

func TestNewUserLastNameHasInvalidCharacters(t *testing.T) {
	_, err := NewUser("John", "Sm1th", "07875166726", "johnsmith@gmail.com")
	if err != ErrNameContainsInvalidCharacters {
		t.Errorf("Expected error: %s", ErrNameTooLong.Error())
	}
}
