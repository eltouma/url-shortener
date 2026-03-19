package models

import (
    "crypto/rand"
    "encoding/base64"
    "log"
    "url-shortener/storage"
)

// generate random short url of length 6 and check if it already exists to avoid collision
func GenerateShortURL(store *storage.Store) string {
    b := make([]byte, 6)
    _, err := rand.Read(b)
    if err != nil {
        log.Fatalf("Failed to generate random url: %v", err)
    }
    shortUrl := base64.URLEncoding.EncodeToString(b)[:6]
    if !store.Exists(shortUrl) {
        return shortUrl
    }
    return GenerateShortURL(store)
}

// Replacing recursion with a for loop: more idiomatic, more robust, simpler to maintain
/*
func GenerateShortURL(store *storage.Store) string {
    for {
        b := make([]byte, 6)
        _, err := rand.Read(b)
        if err != nil {
            log.Fatalf("Failed to generate random characters: %v", err)
        }
        shortURL := base64.URLEncoding.EncodeToString(b)[:6]
        if !store.Exists(shortURL) {
            return shortURL
        }
    }
}
*/
