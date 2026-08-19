package repository

import (
	"context"
	"database/sql"
)

type BibleRepository interface {
	GetVersions(ctx context.Context) ([]Version, error)
	GetTestaments(ctx context.Context) ([]Testament, error)
	GetTestament(ctx context.Context, id int) (*Testament, error)
	GetBooks(ctx context.Context) ([]Book, error)
	GetBook(ctx context.Context, id int) (*Book, error)
	GetTestamentBooks(ctx context.Context, testamentId int) ([]Book, error)
	GetBookVerses(ctx context.Context, bookId int) ([]Verse, error)
	GetChapterVerses(ctx context.Context, bookId int, chapterId int) ([]Verse, error)
	GetVerse(ctx context.Context, bookId int, chapterId int, verseId int) (*Verse, error)
}

type sqliteBibleRepo struct {
	db *sql.DB
}

func NewBibleRepository(db *sql.DB) BibleRepository {
	return &sqliteBibleRepo{db: db}
}

func (r *sqliteBibleRepo) GetBookVerses(ctx context.Context, bookId int) ([]Verse, error) {
	query := `SELECT id, version_code, verse_text, book_id, chapter_number, verse_number FROM verses WHERE book_id = ?;`
	rows, err := r.db.QueryContext(ctx, query, bookId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var verses []Verse
	for rows.Next() {
		var verse Verse
		if err := rows.Scan(&verse.ID, &verse.VersionCode, &verse.VerseText, &verse.BookID, &verse.ChapterNumber, &verse.VerseNumber); err != nil {
			return nil, err
		}
		verses = append(verses, verse)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return verses, nil
}

func (r *sqliteBibleRepo) GetChapterVerses(ctx context.Context, bookId int, chapterId int) ([]Verse, error) {
	query := `SELECT id, version_code, verse_text, book_id, chapter_number, verse_number FROM verses WHERE book_id = ? AND chapter_number = ?;`
	rows, err := r.db.QueryContext(ctx, query, bookId, chapterId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var verses []Verse
	for rows.Next() {
		var verse Verse
		if err := rows.Scan(&verse.ID, &verse.VersionCode, &verse.VerseText, &verse.BookID, &verse.ChapterNumber, &verse.VerseNumber); err != nil {
			return nil, err
		}
		verses = append(verses, verse)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return verses, nil
}

func (r *sqliteBibleRepo) GetVerse(ctx context.Context, bookId int, chapterId int, verseId int) (*Verse, error) {
	query := `SELECT id, version_code, verse_text, book_id, chapter_number, verse_number FROM verses WHERE book_id = ? AND chapter_number = ? AND verse_number = ?;`
	rows, err := r.db.QueryContext(ctx, query, bookId, chapterId, verseId)
	if err != nil {
		return nil, err
	}
	var verse Verse
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&verse.ID, &verse.VersionCode, &verse.VerseText, &verse.BookID, &verse.ChapterNumber, &verse.VerseNumber); err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &verse, nil;
}

func (r *sqliteBibleRepo) GetTestamentBooks(ctx context.Context, testamentId int) ([]Book, error) {
	query := `SELECT id, title, testament_id FROM books WHERE testament_id = ?;`
	rows, err := r.db.QueryContext(ctx, query, testamentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.TestamentID); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func (r *sqliteBibleRepo) GetBook(ctx context.Context, id int) (*Book, error) {
	var book Book;
	query := `SELECT id, title, testament_id FROM books WHERE id = ?;`
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&book.ID, &book.Title, &book.TestamentID); err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &book, nil;
}

func (r *sqliteBibleRepo) GetBooks(ctx context.Context) ([]Book, error) {
	query := `SELECT id, title, testament_id FROM books;`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.TestamentID); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func (r *sqliteBibleRepo) GetTestament(ctx context.Context, id int) (*Testament, error) {
	var testament Testament;
	query := `SELECT id, description FROM testaments WHERE id = ?;`
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&testament.ID, &testament.Description); err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &testament, nil;
}

func (r *sqliteBibleRepo) GetTestaments(ctx context.Context) ([]Testament, error) {
	query := `SELECT id, description FROM testaments;`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var testaments []Testament
	for rows.Next() {
		var testament Testament
		if err := rows.Scan(&testament.ID, &testament.Description); err != nil {
			return nil, err
		}
		testaments = append(testaments, testament)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return testaments, nil
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

