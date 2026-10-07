package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	_ = godotenv.Load()
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		fmt.Println("DATABASE_URL environment variable not set")
		return
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to database:", ErrDatabaseConnectionFailure)
		return
	}
	defer db.Close()

	UserService := NewUserService(NewUserRepo(db))
	AppointmentService := NewAppointmentService(NewAppointmentRepo(db))

	runCLI(UserService, AppointmentService)

}
