package main

import "errors"

var ErrLoginIdMustBePositive = errors.New("login ID must be a positive number")
var ErrNameTooShort = errors.New("name must be at least 1 character")
var ErrNameTooLong = errors.New("name cannot be more than 20 characters")
var ErrNameContainsInvalidCharacters = errors.New("name can only contain A-Z, hyphens, and spaces")
var ErrInvalidPhone = errors.New("phone number must be between 10 and 13 digits")
var ErrInvalidPhoneCharacters = errors.New("phone number can only contain digits 0-9")
var ErrInvalidEmail = errors.New("must be a valid email address")
var ErrDuplicateEmail = errors.New("email already taken")
var ErrEmailNotFound = errors.New("email not found")
var ErrInvalidEmailOrPassword = errors.New("invalid email or password")
