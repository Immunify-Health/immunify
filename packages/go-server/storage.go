package main

import "database/sql"

type Storage interface {
	CreateRecord(*Record) error
	DeleteRecord(int) error
	UpdateRecord(*Record) error
	GetAccountByID(int) (*Record, error)
}

type PostgresStore struct {
	db *sql.DB
}
