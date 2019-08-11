package backend

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

// Infra holds the database connection.
type Infra struct {
	db *sql.DB
}

// InitInfra Initializes the database connection.
func InitInfra(db *sql.DB) *Infra {
	return &Infra{db: db}
}

// Backend stores pointers to the domain models.
type Backend struct {
	ProductRepo *Product
	UserRepo    *User
}

// Init Initializes required connections for our application.
func Init(db *sql.DB) *Backend {
	infra := InitInfra(db)
	return &Backend{
		ProductRepo: &Product{Infra: infra},
		UserRepo:    &User{Infra: infra},
	}
}

// Connect will establish connection with postgres database.
func Connect(host, user, pass, db, port string) (*sql.DB, error) {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, db)
	conn, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, errors.Errorf("Unable to establish connection with postgres: %v", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, errors.Errorf("Unable ping to postgres: %v", err)
	}
	return conn, nil
}
