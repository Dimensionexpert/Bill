package database

import (
	"database/sql"
	_ "embed"

	"fmt"

	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var schema string

func Open(path string) (*sql.DB, error) {

	// sql.open only opens the database connection pool, bt dose not gaurantee the connection to the database

	dsn := fmt.Sprintf("file:%s?_pragma=foreign_keys(1)", path)
	database, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf(":( can't opening the database connection pool :%w", err)
	}

	// for establisng the connection db.ping is used. it checks the state of connection and establish if necessary
	if err := database.Ping(); err != nil {
		database.Close()
		return nil, fmt.Errorf(":( ping the database :%w", err)
	}

	// create the tables by executing the schema.
	if _, err := database.Exec(schema); err != nil {
		database.Close()
		return nil, fmt.Errorf(":( executing the schema :%w", err)

	}

	return database, nil
}
