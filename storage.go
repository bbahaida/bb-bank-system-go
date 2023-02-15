package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(account *Account) error
	DeleteAccount(id int) error
	UpdateAccount(account *Account) error
	GetAccountByID(id int) (*Account, error)
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS ACCOUNT (
			id serial primary key,
			first_name varchar(100),
			last_name varchar(100),
			account_number serial,
			balance serial,
			created_at timestamp
		    )`
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "postgres://postgres:postgres@localhost/bbank?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) CreateAccount(account *Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}
func (s *PostgresStore) UpdateAccount(account *Account) error {
	return nil
}
func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
