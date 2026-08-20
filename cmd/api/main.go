package main

import (
	"log"
	"net/http"
	"time"
	"bibleservice/internal/database"
	"bibleservice/internal/handler"
	"bibleservice/internal/repository"
)

func main() {
	db, err := database.NewSQLite("bible.db")
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer db.Close()
	bibleService := repository.NewBibleService(db)
	restHandler := handler.NewRestHandler(bibleService)
	documentHandler := handler.NewDocumentHandler(bibleService)
	mux := http.NewServeMux()
	restHandler.RegisterRoutes(mux)
	documentHandler.RegisterRoutes(mux)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	log.Println("Server running on port 8080...")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}

