package repository

import (
	"context"
	"database/sql"
)

type BibleService interface {
	GetVersions(ctx context.Context) ([]Version, error)
	GetTestaments(ctx context.Context) ([]Testament, error)
	GetTestament(ctx context.Context, id int) (*Testament, error)
	GetBooks(ctx context.Context) ([]Book, error)
	GetBook(ctx context.Context, id int) (*Book, error)
	GetBookChapterList(ctx context.Context, bookId int) ([]int, error)
	GetTestamentBooks(ctx context.Context, testamentId int) ([]Book, error)
	GetBookVerses(ctx context.Context, bookId int) ([]Verse, error)
	GetChapterVerses(ctx context.Context, bookId int, chapterId int) ([]Verse, error)
	GetVerse(ctx context.Context, bookId int, chapterId int, verseId int) (*Verse, error)
}

type bibleDao struct {
	db *sql.DB
}

func NewBibleService(db *sql.DB) BibleService {
	return &bibleDao{db: db}
}

func (dao *bibleDao) GetBookVerses(ctx context.Context, bookId int) ([]Verse, error) {
	query := `SELECT id, version_code, verse_text, book_id, chapter_number, verse_number FROM verses WHERE book_id = ?;`
	rows, err := dao.db.QueryContext(ctx, query, bookId)
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

func (dao *bibleDao) GetChapterVerses(ctx context.Context, bookId int, chapterId int) ([]Verse, error) {
	query := `SELECT id, version_code, verse_text, book_id, chapter_number, verse_number FROM verses WHERE book_id = ? AND chapter_number = ?;`
	rows, err := dao.db.QueryContext(ctx, query, bookId, chapterId)
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

func (dao *bibleDao) GetVerse(ctx context.Context, bookId int, chapterId int, verseId int) (*Verse, error) {
	query := `SELECT id, version_code, verse_text, book_id, chapter_number, verse_number FROM verses WHERE book_id = ? AND chapter_number = ? AND verse_number = ?;`
	rows, err := dao.db.QueryContext(ctx, query, bookId, chapterId, verseId)
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

func (dao *bibleDao) GetTestamentBooks(ctx context.Context, testamentId int) ([]Book, error) {
	query := `SELECT id, title, testament_id FROM books WHERE testament_id = ?;`
	rows, err := dao.db.QueryContext(ctx, query, testamentId)
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

func (dao *bibleDao) GetBook(ctx context.Context, id int) (*Book, error) {
	var book Book;
	query := `SELECT id, title, testament_id FROM books WHERE id = ?;`
	rows, err := dao.db.QueryContext(ctx, query, id)
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

func (dao *bibleDao) GetBookChapterList(ctx context.Context, bookId int) ([]int, error) {
	query := `SELECT DISTINCT chapter_number FROM verses WHERE book_id = ?;`
	rows, err := dao.db.QueryContext(ctx, query, bookId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var chapters []int
	for rows.Next() {
		var chapter int
		if err := rows.Scan(&chapter); err != nil {
			return nil, err
		}
		chapters = append(chapters, chapter)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return chapters, nil;
}

func (dao *bibleDao) GetBooks(ctx context.Context) ([]Book, error) {
	query := `SELECT id, title, testament_id FROM books;`
	rows, err := dao.db.QueryContext(ctx, query)
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

func (dao *bibleDao) GetTestament(ctx context.Context, id int) (*Testament, error) {
	var testament Testament;
	query := `SELECT id, description FROM testaments WHERE id = ?;`
	rows, err := dao.db.QueryContext(ctx, query, id)
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

func (dao *bibleDao) GetTestaments(ctx context.Context) ([]Testament, error) {
	query := `SELECT id, description FROM testaments;`
	rows, err := dao.db.QueryContext(ctx, query)
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

func (dao *bibleDao) GetVersions(ctx context.Context) ([]Version, error) {
	query := `SELECT version_code, version_name FROM versions;`
	rows, err := dao.db.QueryContext(ctx, query)
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

