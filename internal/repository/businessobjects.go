package repository

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
