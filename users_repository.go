package main

import "database/sql"

type UserRepository interface {
	CreateUser(u user) error
	FindByEmail(email string) (user, error)
	FindByEmailForLogin(email string) (user, []byte, error)
}

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) FindByEmail(email string) (user, error) {
	var u user
	err := r.DB.QueryRow(
		`SELECT id, first_name, last_name, phone, email
		FROM users
		WHERE email = $1`,
		email,
	).Scan(&u.id, &u.firstName, &u.lastName, &u.phone, &u.email)

	if err == sql.ErrNoRows {
		return user{}, ErrEmailNotFound
	}

	if err != nil {
		return user{}, err
	}

	return u, nil

}

func (r *UserRepo) CreateUser(u user) error {

	query := "INSERT INTO users (first_name, last_name, phone, email, password_hash) VALUES ($1, $2, $3, $4, $5)"

	_, err := r.DB.Exec(query, u.firstName, u.lastName, u.phone, u.email, u.passwordHash)

	// let service handle the err!
	return err
}

func (r *UserRepo) FindByEmailForLogin(email string) (user, []byte, error) {
	var passwordHash []byte
	var u user
	err := r.DB.QueryRow(
		`SELECT id, first_name, last_name, phone, email, password_hash
		FROM users
		WHERE email = $1`,
		email,
	).Scan(&u.id, &u.firstName, &u.lastName, &u.phone, &u.email, &passwordHash)

	if err == sql.ErrNoRows {
		return u, passwordHash, ErrInvalidEmailOrPassword
	}

	if err != nil {
		return user{}, []byte{}, err
	}

	return u, passwordHash, nil
}
