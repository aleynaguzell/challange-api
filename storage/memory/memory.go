package memory

import (
	"errors"
	"sync"
)

type Store struct {
	mutex *sync.Mutex
	data map[string]string
}

func (s *Store) Store(key,value string) error {
	if len(key) == 0 || len(value) == 0 {
		return errors.New("key or value can not be empty")
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()


	s.data[key] = value

	return nil
}

func (s *Store) Get(key string) (string, error) {
	if len(key) == 0  {
		return "",errors.New("key can not be empty")
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if val, ok := s.data[key]; ok {
		return val, nil
	}

	 return "",errors.New("key is not found")
}

// New creates in memory store
func New() *Store {
	return &Store{
		data: make(map[string]string),
		mutex:   &sync.Mutex{},
	}
}
