package database

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	DB *sqlx.DB
}

func NewDatabase(
	username,
	password,
	host,
	databaseName string,
	port int,
) (*Database, error) {
	params := []string{
		"parseTime=true",
		"timeout=30s",      // Note: this is the connection timeout, not the query timeout
		"writeTimeout=30s", // Note: this the timeout for awaiting a block to be written, not the entire query.
		"readTimeout=30s",  // Note: this the timeout for awaiting a block to be read, not the entire result set.
	}
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", username, password, host, port, databaseName, strings.Join(params, "&"))
	db, err := open(uri)
	if err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil

}

func (d *Database) Close() error {
	return d.DB.Close()
}

func open(uri string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", uri)

	if err != nil {
		return nil, err
	}

	return db, nil
}
