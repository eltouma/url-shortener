package handlers

import (
    "fmt"
    "net/http"
    "url-shortener/models"
    "url-shortener/storage"
    "os"
//    "regexp"
)

type Handler struct {
    store *storage.Store
}

func NewHandler(store *storage.Store) *Handler {
    return &Handler{store: store}
}

var siteUrl = os.Getenv("SITEURL")

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
    if siteUrl == "" {
        siteUrl = "http://localhost:8081/"
    }
    fmt.Println(r)
    fmt.Println(r.Method)
    // à modifier, voir le repo 1
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed...", http.StatusMethodNotAllowed)
        return
    }
    w.Header().Set("Content-Type", "application/json")

    r.ParseForm()
    fmt.Println(r)
    longURL := r.FormValue("url")
    fmt.Println("\n")
    fmt.Println(longURL)
    if longURL == "" {
        http.Error(w, "missing 'url' parameter", http.StatusBadRequest)
        return
    }

    shortURL := models.GenerateShortURL(h.store)
    h.store.Save(shortURL, longURL)

    w.WriteHeader(http.StatusCreated)

    // carriage return is necessary otherwise shortURL ends by % and redirect to a 404
    w.Write([]byte(siteUrl + shortURL + "\n"))
}

func (h *Handler) RedirectURL(w http.ResponseWriter, r *http.Request) {
    /*
    getPage := false
    if r.URL.Path == "/get" {
        getPage = true
    }
    */
    if r.Method != http.MethodGet {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    shortURL := r.URL.Path[1:]
    fmt.Println("shotURL:", shortURL)

    longURL, exists := h.store.Get(shortURL)
    if !exists {
        http.NotFound(w, r)
        return
    }

    http.Redirect(w, r, longURL, http.StatusFound)
}
