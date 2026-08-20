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
	mux.HandleFunc("GET /", h.ServeHomePage)
	mux.HandleFunc("GET /books", h.ServeBooksPage)
	mux.HandleFunc("GET /bookList", h.GetBookList)
	mux.HandleFunc("GET /book/{id}", h.ServeBookDetailPage)
}

func (h *DocumentHandler) ServeHomePage(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "index.html", nil)
}

func (h *DocumentHandler) ServeBooksPage(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "books.html", nil)
}

func (h *DocumentHandler) GetBookList(w http.ResponseWriter, r *http.Request) {
	books, err := h.bibleService.GetBooks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.Header.Get("HX-Request") != "true" {
		http.NotFound(w, r)
		return
	}
	for _, book := range books {
		h.templates.ExecuteTemplate(w, "booklist.html", book)
	}
}

func (h *DocumentHandler) ServeBookDetailPage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	bookId, _ := strconv.Atoi(id)
	h.templates.ExecuteTemplate(w, "bookdetail.html", bookId)
}

