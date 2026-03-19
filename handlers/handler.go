package handlers

import (
    "net/http"
    "net/url"
    "log"
    "url-shortener/models"
    "url-shortener/storage"
)

type Handler struct {
    store *storage.Store
    siteUrl string
    infoLogger *log.Logger
}

func NewHandler(store *storage.Store, siteUrl string, infoLogger *log.Logger) *Handler {
    return &Handler{
        store: store,
        siteUrl: siteUrl,
        infoLogger: infoLogger,
    }
}

func isValidURL(longUrl string) bool {
    parsedUrl, err := url.ParseRequestURI(longUrl)
    if err != nil {
        return false
    }
    if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
        return false
    }
    if parsedUrl.Host == "" {
        return false
    }
    return true
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Error parsing form", http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "text/plain")
    longUrl := r.FormValue("url")
    if longUrl == "" {
        http.Error(w, "Missing 'url' parameter", http.StatusBadRequest)
        return
    }
    if !isValidURL(longUrl) {
        http.Error(w, "Invalid 'url' parameter", http.StatusBadRequest)
        return
    }

    shortUrl := models.GenerateShortURL(h.store)
    h.store.Save(shortUrl, longUrl)

    w.WriteHeader(http.StatusCreated)
    // add a new line for terminal output when using curl
    w.Write([]byte(h.siteUrl + shortUrl + "\n"))
}

func (h *Handler) RedirectURL(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    if r.URL.Path == "/" {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Home page\n"))
        h.infoLogger.Println("GET /")
        return
    }

    shortUrl := r.URL.Path[1:]
    longUrl, exists := h.store.Get(shortUrl)
    if !exists {
        http.NotFound(w, r)
        return
    }
    h.infoLogger.Println("GET", h.siteUrl + shortUrl, "redirect to ", longUrl)
    http.Redirect(w, r, longUrl, http.StatusFound)
}
