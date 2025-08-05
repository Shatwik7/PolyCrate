package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB is a struct that holds the database connection
type DB struct {
	Conn *sql.DB
}

// NewDB initializes a new database connection
func NewDB(dataSourceName string) (*DB, error) {
	conn, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// schemaBytes, err := os.ReadFile("./schema.sql") // read schema file
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to read schema.sql: %w", err)
	// }

	// schema := string(schemaBytes)
	// queries := strings.Split(schema, ";")

	// // Execute each statement
	// for _, query := range queries {
	// 	query = strings.TrimSpace(query)
	// 	if query == "" {
	// 		continue
	// 	}
	// 	if _, err := conn.Exec(query); err != nil {
	// 		return nil, fmt.Errorf("failed to execute schema statement: %w\nQuery: %s", err, query)
	// 	}
	// }

	return &DB{Conn: conn}, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	if db.Conn != nil {
		if err := db.Conn.Close(); err != nil {
			return fmt.Errorf("failed to close database connection: %w", err)
		}
	}
	return nil
}

// Ping checks if the database connection is alive
func (db *DB) Ping() error {
	if db.Conn == nil {
		return fmt.Errorf("database connection is nil")
	}
	if err := db.Conn.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}
	return nil
}

// Exec executes a query without returning any rows
func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	if db.Conn == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	result, err := db.Conn.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return result, nil
}

// QueryRow executes a query that is expected to return a single row
func (db *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	if db.Conn == nil {
		return nil
	}
	return db.Conn.QueryRow(query, args...)
}

// Query executes a query that returns multiple rows
func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if db.Conn == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	rows, err := db.Conn.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return rows, nil
}

// Begin starts a new transaction
func (db *DB) Begin() (*sql.Tx, error) {
	if db.Conn == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	tx, err := db.Conn.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	return tx, nil
}

// Commit commits a transaction
func (db *DB) Commit(tx *sql.Tx) error {
	if tx == nil {
		return fmt.Errorf("transaction is nil")
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

// Rollback rolls back a transaction
func (db *DB) Rollback(tx *sql.Tx) error {
	if tx == nil {
		return fmt.Errorf("transaction is nil")
	}
	if err := tx.Rollback(); err != nil {
		return fmt.Errorf("failed to rollback transaction: %w", err)
	}
	return nil
}
