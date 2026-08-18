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
	INSERT OR IGNORE INTO versions(version_code, version_name) VALUES('ESV', 'English Standard Version');

	CREATE TABLE IF NOT EXISTS testaments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT,
		UNIQUE(description)
	);
	INSERT OR IGNORE INTO testaments(description) VALUES('Old Testament');
	INSERT OR IGNORE INTO testaments(description) VALUES('New Testament');

	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		testament_id INTEGER NOT NULL REFERENCES testaments(id),
		UNIQUE(title, testament_id)
	);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Genesis', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Exodus', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Leviticus', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Numbers', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Deuteronomy', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Joshua', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Judges', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Ruth', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('1 Samuel', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('2 Samuel', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('1 Kings', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('2 Kings', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('1 Chronicles', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('2 Chronicles', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Ezra', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Nehemiah', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Esther', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Job', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Psalm', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Proverbs', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Ecclesiastes', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Song of Solomon', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Isaiah', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Jeremiah', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Lamentations', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Ezekiel', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Daniel', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Hosea', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Joel', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Amos', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Obadiah', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Jonah', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Micah', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Nahum', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Habakkuk', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Zephaniah', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Haggai', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Zechariah', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Malachi', 1);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Matthew', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Mark', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Luke', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('John', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Acts', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Romans', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('1 Corinthians', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('2 Corinthians', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Galatians', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Ephesians', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Philippians', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Colossians', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('1 Thessalonians', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('2 Thessalonians', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('1 Timothy', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('2 Timothy', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Titus', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Philemon', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Hebrews', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('James', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('1 Peter', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('2 Peter', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('1 John', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('2 John', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('3 John', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Jude', 2);
	INSERT OR IGNORE INTO books(title, testament_id) VALUES('Revelation', 2);

	CREATE TABLE IF NOT EXISTS verses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		version_code TEXT NOT NULL REFERENCES versions(version_code),
		verse_text TEXT NOT NULL,
		book_id INTEGER NOT NULL REFERENCES books(id),
		chapter_number INTEGER NOT NULL,
		verse_number INTEGER NOT NULL,
		UNIQUE(version_code, book_id, chapter_number, verse_number)
		
	);
	`
	_, err := db.Exec(query)
	return err
}


