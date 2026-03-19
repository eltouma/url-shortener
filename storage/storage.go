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

func (s *Store) Save(shortUrl, longUrl string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.urls[shortUrl] = longUrl
}

func (s *Store) Get(shortUrl string) (string, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    longUrl, exists := s.urls[shortUrl]
    return longUrl, exists
}

// verify if shortUrl already exists in memory
func (s *Store) Exists(shortUrl string) bool {
    s.mu.RLock()
    defer s.mu.RUnlock()
    _, ok := s.urls[shortUrl]
    return ok
}
