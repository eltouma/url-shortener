package storage

import (
    "sync"
)

// in-memory storage: key: short URL, value: long URL
// using mutex for concurrency: read/write
type Store struct {
    urls map[string]string
    mu sync.RWMutex
}

func NewStore() *Store {
    return &Store{
        urls:make(map[string]string),
    }
}

func (s *Store) Save(shortURL, longURL string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.urls[shortURL] = longURL
}

func (s *Store) Get(shortURL string) (string, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    longURL, exists := s.urls[shortURL]
    return longURL, exists
}

// verify if shortURL already exists in memory
func (s *Store) Exists(shortURL string) bool {
    s.mu.RLock()
    defer s.mu.RUnlock()
    _, ok := s.urls[shortURL]
    return ok
}
