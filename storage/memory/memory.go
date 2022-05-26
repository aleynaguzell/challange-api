package memory

import (
	"errors"
	"sync"
)

type Store struct {
	m    *sync.Mutex
	data map[string]string
}

// New creates inmemory store
func New() *Store {
	return &Store{
		data: make(map[string]string),
		m:    &sync.Mutex{},
	}
}

func (s *Store) Store(key, value string) error {
	if len(key) == 0 || len(value) == 0 {
		return errors.New("key or value can not be empty")
	}
	s.m.Lock()
	defer s.m.Unlock()

	s.data[key] = value

	return nil
}

func (s *Store) Get(key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("key can not be empty")
	}
	s.m.Lock()
	defer s.m.Unlock()
	if val, ok := s.data[key]; ok {
		return val, nil
	}

	return "", errors.New("key is not found")
}
