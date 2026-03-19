package main

import (
    "net/http"
    "url-shortener/handlers"
    "url-shortener/storage"
    "log"
    "os"
)

var (
    infoLogger  = log.New(os.Stdout, "[info] ", log.LstdFlags)
    errorLogger = log.New(os.Stderr, "[error] ", log.LstdFlags|log.Lshortfile)
)

var siteUrl = os.Getenv("SITEURL")

func main() {
    if siteUrl == "" {
        siteUrl = "http://localhost:8081/"
    }
    store := storage.NewStore()
    handler := handlers.NewHandler(store, siteUrl, infoLogger)

    http.HandleFunc("GET /", handler.RedirectURL)
    http.HandleFunc("POST /shorten", handler.ShortenURL)
    port := ":8081"
    infoLogger.Println("Server is running on port:", port)

    if err := http.ListenAndServe(port, nil); err != nil {
        errorLogger.Fatalf("Failed to start server on %s: %v", port, err)
    }
}
