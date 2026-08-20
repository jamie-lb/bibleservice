package handler

import (
	"html/template"
	"net/http"
	"strconv"
)

type DocumentHandler struct {
	templates *template.Template
}

func NewDocumentHandler() *DocumentHandler {
	return &DocumentHandler{templates: template.Must(template.ParseGlob("internal/handler/templates/*.html"))}
}

func (h *DocumentHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /", h.ServeHomePage)
	mux.HandleFunc("GET /books", h.ServeBooksPage)
	mux.HandleFunc("GET /book/{id}", h.ServeBookDetailPage)
}

func (h *DocumentHandler) ServeHomePage(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "index.html", nil)
}

func (h *DocumentHandler) ServeBooksPage(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "books.html", nil)
}

func (h *DocumentHandler) ServeBookDetailPage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	bookId, _ := strconv.Atoi(id)
	h.templates.ExecuteTemplate(w, "bookdetail.html", bookId)
}

