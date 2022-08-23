package repository

import "database/sql"

type Repository interface {
	Connect()
	Disconnect()
	GetDB() *sql.DB
}
