package postgresql

import (
	"database/sql"
	"payment/RegistrationService/internal/port"
)

type Database struct {
	db *sql.DB
}

func NewRepository(d *sql.DB) port.DatabaseRepository {
	return &Database{
		db: d,
	}
}

func (d *Database) Register() error {
	return nil
}