package rw_mutex_cache

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("value not found")

type Cache interface {
	Set(key, value string) error
	Get(Key string) (string, error)
	Delete(Key string) error
}

type MutexStorage struct {
	m  map[string]string
	mu *sync.RWMutex
}

func New() *MutexStorage {
	return &MutexStorage{
		m:  make(map[string]string),
		mu: new(sync.RWMutex),
	}
}

func (s *MutexStorage) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *MutexStorage) Get(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	value, ok := s.m[key]
	return value, ok
}

func (s *MutexStorage) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, key)
}
