package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

var (
	ErrInvalidTransaction = errors.New("no valid transaction")
)

// START1 OMIT
type database struct {
	db *sql.DB
}

type IDatabase interface {
	Commit() error
	Rollback() error
	Sql() *sql.DB

	MustBeginTransaction() IDatabase // HL
}

// STOP1 OMIT

type Config struct {
	IP       string
	User     string
	Password string
	Name     string
}

func NewDatabase(c Config) IDatabase {
	var err error

	database := new(database)

	database.db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", c.User, c.Password, c.IP, c.Name))
	if err != nil {
		log.Printf("db: %s", err.Error())
	}

	if connSuccess := database.db.Ping(); !connSuccess {
		panic("Could not ping database")
	}

	return database
}

func (d *database) Sql() *sql.DB {
	return d.db
}

// MustGetTransaction returns new database object with opened transaction.
// Panics if transaction cannot be opened
func (d *database) MustBeginTransaction() IDatabase {
	tx, err := d.db.Begin()

	if err != nil {
		panic(err.Error())
	}

	return &database{
		db: tx,
	}
}

// START2 OMIT
func (d *database) Commit() error {
	if tx, ok := d.db.(*sql.Tx); ok { // HL
		return tx.Commit()
	} else {
		return ErrInvalidTransaction
	}

	return nil
}

// STOP2 OMIT

func (d *database) Rollback() error {
	if tx, ok := d.db.(*sql.Tx); ok { // HL
		return tx.Rollback()
	} else {
		return ErrInvalidTransaction
	}

	return nil
}
