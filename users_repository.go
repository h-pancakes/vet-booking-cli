package main

import "database/sql"

type UserRepository interface {
	Create(u user) (user, error)
	FindByEmail(email string) (user, error)
}

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) FindUserByEmail(email string) (bool, error) {
	var exists bool

	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)"

	err := r.DB.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
