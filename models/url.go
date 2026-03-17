package models

import (
    "crypto/rand"
    "encoding/base64"
    "log"
    "url-shortener/storage"
)

// generate random short url and check if it already exists to avoid collision
func GenerateShortURL(store *storage.Store) string {
    b := make([]byte, 6)
    _, err := rand.Read(b)
    if err != nil {
        log.Fatalf("Failed to generate random characters: %v", err)
    }
    shortURL := base64.URLEncoding.EncodeToString(b)[:6]
    if !store.Exists(shortURL) {
        return shortURL
    }
    return GenerateShortURL(store)
}

// Remplacer la récursivité par un for : plus idiomatique, plus robuste, plus simple à maintenir
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
