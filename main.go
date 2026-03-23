package main

import (
    "net"
    "net/http"
    "net/url"
    "url-shortener/handlers"
    "url-shortener/storage"
    "strings"
    "log"
    "os"
)

var (
    infoLogger  = log.New(os.Stdout, "[info] ", log.LstdFlags)
    errorLogger = log.New(os.Stderr, "[error] ", log.LstdFlags|log.Lshortfile)
)

var siteUrl = os.Getenv("SITEURL")
var defaultUrl = "http://localhost:8081/"

func verifySiteUrl(siteUrl string) string {
    if siteUrl == "" {
        return defaultUrl
    }
    u, err := url.Parse(siteUrl)
    if err != nil {
        return defaultUrl
    }

    host := u.Hostname()
    port := u.Port()

    if host == "" {
        return defaultUrl
    }

    if port == "" {
        port = "8081"
    }
    u.Host = net.JoinHostPort(host, port)
    validUrl := u.String()
    if !strings.HasSuffix(validUrl, "/") {
        validUrl += "/"
    }
    return validUrl
}

func definePort(siteUrl string) string {
    u, err := url.Parse(siteUrl)
    if err != nil || u.Port() == "" {
        return ":8081"
    }
    return ":" + u.Port()
}

func main() {
    siteUrl = verifySiteUrl(siteUrl)
    // Add '/' at the end of SITEURL if it's missing
    // Otherwhise POST /shorten response is not working
    if !strings.HasSuffix(siteUrl, "/") {
        siteUrl += "/"
    }
    store := storage.NewStore()
    handler := handlers.NewHandler(store, siteUrl, infoLogger)

    http.HandleFunc("GET /", handler.RedirectURL)
    http.HandleFunc("POST /shorten", handler.ShortenURL)
    port := definePort(siteUrl)
    infoLogger.Println("Server is running on port:", port)

    if err := http.ListenAndServe(port, nil); err != nil {
        errorLogger.Fatalf("Failed to start server on %s: %v", port, err)
    }
}
