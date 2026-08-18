package database

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

func NewSQLite(dbPath string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s?_pragma=journal_mode(WAL)&_pragma=foreign_keys(ON)", dbPath)
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	if err := migrate(db); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}
	return db, nil
}

func migrate(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS versions (
		version_code TEXT PRIMARY KEY,
		version_name TEXT NOT NULL
	);
	INSERT INTO versions(version_code, version_name) VALUES('ESV', 'English Standard Version');
	`
	_, err := db.Exec(query)
	return err
}


