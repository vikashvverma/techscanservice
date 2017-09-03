package repository

import (
	"database/sql"
	"fmt"
)

const unknownID = -1

type Execer interface {
	Exec(query string, args ...interface{}) (rowsAffected int64, err error)
	Query(query string, scanner func(rows *sql.Rows) (interface{}, error), args ...interface{}) (interface{}, error)
}

type DB struct {
	conn *sql.DB
}

func New(db *sql.DB) *DB {
	return &DB{conn: db}
}

func (db *DB) Exec(query string, args ...interface{}) (int64, error) {
	result, err := db.conn.Exec(query, args...)
	if err != nil {
		return unknownID, fmt.Errorf("error executing query: %s", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return unknownID, fmt.Errorf("error getting rows affected: %s", err)
	}
	return rowsAffected, nil
}

func (db *DB) Query(query string, scanner func(rows *sql.Rows) (interface{}, error), args ...interface{}) (interface{}, error) {
	rows, err := db.conn.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %s", err)
	}

	return scanner(rows)
}
