package storages

import "fmt"

type SimpleStorage struct {
	Start uint64
	Urls  map[uint64]string
}

func CreateSimpleStorage(startID uint64) *SimpleStorage {
	return &SimpleStorage{
		Start: startID,
		Urls:  make(map[uint64]string),
	}
}

func (s *SimpleStorage) Get(id uint64) (string, error) {
	if url, ok := s.Urls[id]; ok {
		return url, nil
	}

	return "", fmt.Errorf("id %d have not found", id)
}

func (s *SimpleStorage) Put(url string) (uint64, error) {
	id := s.Start
	s.Start++

	s.Urls[id] = url

	return id, nil
}
