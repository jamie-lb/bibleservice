package repository

import (
	"context"
	"database/sql"
)

type Version struct {
	VersionCode string `json:"version_code"`
	VersionName string `json:"version_name"`
}

type Testament struct {
	ID int `json:"id"`
	Description string `json:"description"`
}

type Book struct {
	ID int `json:"id"`
	Title string `json:"description"`
	TestamentID int `json:"testament_id"`
}

type Verse struct {
	ID int `json:"id"`
	VersionCode string `json:"version_code"`
	VerseText string `json:"verse_text"`
	BookID int `json:"book_id"`
	ChapterNumber int `json:"chapter_number"`
	VerseNumber int `json:"verse_number"`
}

type BibleRepository interface {
	GetVersions(ctx context.Context) ([]Version, error)
}

type sqliteBibleRepo struct {
	db *sql.DB
}

func NewBibleRepository(db *sql.DB) BibleRepository {
	return &sqliteBibleRepo{db: db}
}

func (r *sqliteBibleRepo) GetVersions(ctx context.Context) ([]Version, error) {
	query := `SELECT version_code, version_name FROM versions;`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var versions []Version
	for rows.Next() {
		var version Version
		if err := rows.Scan(&version.VersionCode, &version.VersionName); err != nil {
			return nil, err
		}
		versions = append(versions, version)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return versions, nil
}

