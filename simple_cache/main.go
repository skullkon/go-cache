package simple_cache

import "errors"

var ErrNotFound = errors.New("value not found")

type Cache interface {
	Set(key, value string) error
	Get(Key string) (string, error)
	Delete(Key string) error
}

type Storage struct {
	storage map[string]string
}

func New() *Storage {
	return &Storage{storage: make(map[string]string)}
}

func (s *Storage) Set(key, value string) error {
	s.storage[key] = value

	return nil
}

func (s *Storage) Get(key string) (string, error) {
	value, ok := s.storage[key]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (s *Storage) Delete(key string) error {
	delete(s.storage, key)
	return nil
}
