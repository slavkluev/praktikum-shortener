package app

import (
	"fmt"
)

type Shortener struct {
	start uint64
	urls  map[uint64]string
}

func CreateShortener(startID uint64) Shortener {
	return Shortener{
		start: startID,
		urls:  make(map[uint64]string),
	}
}

func (s *Shortener) GetOriginURL(id uint64) (string, error) {
	if url, ok := s.urls[id]; ok {
		return url, nil
	}

	return "", fmt.Errorf("id %d have not found", id)
}

func (s *Shortener) ShortenURL(url string) (uint64, error) {
	id := s.start
	s.start++

	s.urls[id] = url

	return id, nil
}
