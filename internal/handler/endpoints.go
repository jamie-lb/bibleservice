package handler

import (
	"bibleservice/internal/repository"
	"encoding/json"
	"net/http"
	"strconv"
)

type BibleHandler struct {
	repo repository.BibleRepository
}

func NewBibleServiceHandler(repo repository.BibleRepository) *BibleHandler {
	return &BibleHandler{repo: repo}
}

func (h *BibleHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /versions", h.GetVersions)
	mux.HandleFunc("GET /testaments", h.GetTestaments)
	mux.HandleFunc("GET /testament/{id}", h.GetTestament)
	mux.HandleFunc("GET /books", h.GetBooks)
	mux.HandleFunc("GET /book/{id}", h.GetBook)
	mux.HandleFunc("GET /testamentBooks/{testamentId}", h.GetTestamentBooks)
}

func (h *BibleHandler) GetTestamentBooks(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("testamentId")
	testamentId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	books, err := h.repo.GetTestamentBooks(r.Context(), testamentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BibleHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	bookId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	book, err := h.repo.GetBook(r.Context(), bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&book)
}

func (h *BibleHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.repo.GetBooks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BibleHandler) GetTestament(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	testamentId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	testament, err := h.repo.GetTestament(r.Context(), testamentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&testament)
}

func (h *BibleHandler) GetTestaments(w http.ResponseWriter, r *http.Request) {
	testaments, err := h.repo.GetTestaments(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(testaments)
}

func (h *BibleHandler) GetVersions(w http.ResponseWriter, r *http.Request) {
	versions, err := h.repo.GetVersions(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(versions)
}

