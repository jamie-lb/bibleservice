package handler

import (
	"encoding/json"
	"bibleservice/internal/repository"
	"net/http"
)

type BibleHandler struct {
	repo repository.BibleRepository
}

func NewBibleServiceHandler(repo repository.BibleRepository) *BibleHandler {
	return &BibleHandler{repo: repo}
}

func (h *BibleHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /versions", h.GetVersions)
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

