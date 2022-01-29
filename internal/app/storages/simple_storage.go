package storages

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type SimpleStorage struct {
	Start  uint64
	Urls   map[uint64]string
	File   *os.File
	ticker *time.Ticker
	done   chan bool
}

type Record struct {
	ID  uint64
	URL string
}

func CreateSimpleStorage(filename string, syncTime int) (*SimpleStorage, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	if err != nil {
		return nil, err
	}

	lastID, urls, err := loadDataFromFile(file)
	if err != nil {
		return nil, err
	}

	ticker := time.NewTicker(time.Duration(syncTime) * time.Minute)
	done := make(chan bool)
	simpleStorage := &SimpleStorage{
		Start:  lastID + 1,
		Urls:   urls,
		File:   file,
		ticker: ticker,
		done:   done,
	}

	go simpleStorage.synchronize()

	return simpleStorage, nil
}

func loadDataFromFile(file *os.File) (uint64, map[uint64]string, error) {
	var lastID uint64 = 0
	var urls = make(map[uint64]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Bytes()

		record := Record{}
		err := json.Unmarshal(data, &record)
		if err != nil {
			return 0, nil, err
		}

		if record.ID > lastID {
			lastID = record.ID
		}

		urls[record.ID] = record.URL
	}

	return lastID, urls, nil
}

func (s *SimpleStorage) synchronize() {
	for {
		select {
		case <-s.done:
			return
		case <-s.ticker.C:
			err := s.updateDataFile()
			if err != nil {
				return
			}
		}
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

func (s *SimpleStorage) Close() error {
	s.ticker.Stop()
	s.done <- true
	err := s.updateDataFile()

	if err != nil {
		return err
	}

	return s.File.Close()
}

func (s *SimpleStorage) updateDataFile() error {
	err := s.File.Truncate(0)
	if err != nil {
		return err
	}

	_, err = s.File.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	for ID, URL := range s.Urls {
		err := s.writeRecordToFile(Record{
			ID:  ID,
			URL: URL,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *SimpleStorage) writeRecordToFile(record Record) error {
	data, err := json.Marshal(record)

	if err != nil {
		return err
	}

	data = append(data, '\n')
	_, err = s.File.Write(data)

	return err
}
