package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"bibleservice/internal/repository"
)

type DocumentHandler struct {
	templates *template.Template
	bibleService repository.BibleService
}

func NewDocumentHandler(bibleService repository.BibleService) *DocumentHandler {
	return &DocumentHandler{bibleService: bibleService, templates: template.Must(template.ParseGlob("internal/handler/templates/*.html"))}
}

func (h *DocumentHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("internal/handler/templates/css"))))
	mux.HandleFunc("GET /{$}", h.ServeHomePage)
	mux.HandleFunc("GET /books", h.ServeBooksPage)
	mux.HandleFunc("GET /book/{id}", h.ServeBookDetailPage)
	mux.HandleFunc("GET /book/{bookId}/chapter/{chapterId}", h.ServeBookChapterPage)
}

func (h *DocumentHandler) ServeHomePage(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "index.html", nil)
}

func (h *DocumentHandler) ServeBooksPage(w http.ResponseWriter, r *http.Request) {
	books, err := h.bibleService.GetBooks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.templates.ExecuteTemplate(w, "books.html", books)
}

func (h *DocumentHandler) ServeBookDetailPage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	bookId, _ := strconv.Atoi(id)
	book, err := h.bibleService.GetBook(r.Context(), bookId)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	chapters, err := h.bibleService.GetBookChapterList(r.Context(), bookId)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var viewData repository.BookViewData = repository.BookViewData{Book: *book, Chapters: chapters}
	h.templates.ExecuteTemplate(w, "bookdetail.html", viewData)
}

func (h *DocumentHandler) ServeBookChapterPage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("bookId")
	chapterNumber := r.PathValue("chapterId")
	bookId, _ := strconv.Atoi(id)
	chapterId, _ := strconv.Atoi(chapterNumber)
	book, err := h.bibleService.GetBook(r.Context(), bookId)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	verses, err := h.bibleService.GetChapterVerses(r.Context(), bookId, chapterId)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var viewData repository.BookChapterViewData = repository.BookChapterViewData{Book: *book, ChapterNumber: chapterId, Verses: verses}
	h.templates.ExecuteTemplate(w, "bookchapter.html", viewData)
}

