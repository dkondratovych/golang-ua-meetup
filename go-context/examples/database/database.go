package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// ErrInvalidTransaction invalid transaction when you are trying to `Commit` or `Rollback`
	ErrInvalidTransaction = errors.New("no valid transaction")

	// ErrCantStartTransaction can't start transaction when you are trying to start one with `Begin`
	ErrCantStartTransaction = errors.New("can't start transaction")
)

// START 1A OMIT
type sqlCommon interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type sqlDb interface {
	Begin() (*sql.Tx, error)
	MustBegin() *sql.Tx
}

type sqlTx interface {
	sqlCommon
	Commit() error
	Rollback() error
}

// STOP 1A OMIT

// START1 OMIT
type database struct {
	db sqlCommon
}

type Database interface {
	Sql() *sql.DB

	Commit() error
	Rollback() error
	PingDB() error

	MustBeginTransaction() Database // HL
}

// STOP1 OMIT

type Config struct {
	IP       string
	User     string
	Password string
	Name     string
}

func NewDatabase(c Config) (Database, error) {
	var err error

	database := new(database)

	database.db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", c.User, c.Password, c.IP, c.Name))
	if err != nil {
		return nil, err
	}

	return database, nil
}

func (d *database) Sql() *sql.DB {
	return d.db.(*sql.DB)
}

// MustBeginTransaction returns new database object with opened transaction.
// Panics if transaction cannot be opened
func (d *database) MustBeginTransaction() Database {
	var db sqlDb
	var tx sqlTx

	db, ok := d.db.(sqlDb)
	if ok {
		tx = db.MustBegin()
	} else {
		panic(ErrCantStartTransaction)
	}

	return &database{
		db: tx,
	}
}

// START2 OMIT
func (d *database) Commit() error {
	if tx, ok := d.db.(sqlTx); ok { // HL
		return tx.Commit()
	}

	return ErrInvalidTransaction
}

// STOP2 OMIT
func (d *database) Rollback() error {
	if tx, ok := d.db.(sqlTx); ok { // HL
		return tx.Rollback()
	}

	return ErrInvalidTransaction
}

func (d *database) PingDB() error {
	return d.Sql().Ping()
}
