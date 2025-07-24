package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
}

func (db *DB) LoadSQLSchema(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(content))
	return err
}

func NewDB(user, password, host, dbname string, port int) (*DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}
	DB := &DB{Conn: conn}

	if err = DB.LoadSQLSchema("./schema.sql"); err != nil {
		log.Fatal("Failed to load schema:", err)
		return nil, err
	}

	return DB, nil
}

func (db *DB) Close() {
	db.Conn.Close()
}

func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.Conn.Exec(query, args...)
}

func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.Conn.Query(query, args...)
}

func (db *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.Conn.QueryRow(query, args...)
}
