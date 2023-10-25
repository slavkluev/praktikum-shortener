package storages

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// SimpleStorage хранилище, использующее файл
type SimpleStorage struct {
	Start   uint64
	Records map[uint64]Record
	File    *os.File
	ticker  *time.Ticker
	done    chan bool
	mu      sync.Mutex
}

// CreateSimpleStorage создание SimpleStorage
func CreateSimpleStorage(filename string, syncTime int) (*SimpleStorage, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	if err != nil {
		return nil, err
	}

	lastID, records, err := loadDataFromFile(file)
	if err != nil {
		return nil, err
	}

	ticker := time.NewTicker(time.Duration(syncTime) * time.Second)
	done := make(chan bool)
	simpleStorage := &SimpleStorage{
		Start:   lastID + 1,
		Records: records,
		File:    file,
		ticker:  ticker,
		done:    done,
	}

	go simpleStorage.synchronize()

	return simpleStorage, nil
}

func loadDataFromFile(file *os.File) (uint64, map[uint64]Record, error) {
	var lastID uint64 = 0
	var urls = make(map[uint64]Record)
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

		urls[record.ID] = record
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

// Get получение записи по ID
func (s *SimpleStorage) Get(ctx context.Context, id uint64) (Record, error) {
	if record, ok := s.Records[id]; ok {
		return record, nil
	}

	return Record{}, fmt.Errorf("id %d have not found", id)
}

// GetByOriginURL получение записи по оригинальному URL
func (s *SimpleStorage) GetByOriginURL(ctx context.Context, originURL string) (Record, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, record := range s.Records {
		if record.URL == originURL && !record.Deleted {
			return record, nil
		}
	}

	return Record{}, fmt.Errorf("originURL %s have not found", originURL)
}

// GetByUser получение всех записей у пользователя
func (s *SimpleStorage) GetByUser(ctx context.Context, userID string) ([]Record, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	records := make([]Record, 0)

	for _, record := range s.Records {
		if record.User == userID && !record.Deleted {
			records = append(records, record)
		}
	}

	return records, nil
}

// Put вставка одной записи
func (s *SimpleStorage) Put(ctx context.Context, record Record) (uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	record.ID = s.Start
	s.Start++

	s.Records[record.ID] = record

	return record.ID, nil
}

// Close закрытие файлового хранилища
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
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.File.Truncate(0)
	if err != nil {
		return err
	}

	_, err = s.File.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	buf := bytes.Buffer{}
	for _, record := range s.Records {
		data, err := json.Marshal(record)

		if err != nil {
			return err
		}

		buf.Write(data)
		buf.WriteRune('\n')
	}

	_, err = s.File.Write(buf.Bytes())

	return err
}

// PutRecords вставка множества записей
func (s *SimpleStorage) PutRecords(ctx context.Context, records []BatchRecord) ([]BatchRecord, error) {
	return nil, fmt.Errorf("method has not implemented")
}

// Ping проверка доступности хранилища
func (s *SimpleStorage) Ping(ctx context.Context) error {
	return fmt.Errorf("method has not implemented")
}

// DeleteRecords удаление нескольких записей
func (s *SimpleStorage) DeleteRecords(ctx context.Context, ids []uint64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, id := range ids {
		if record, ok := s.Records[id]; ok {
			record.Deleted = true
			s.Records[record.ID] = record
		} else {
			return fmt.Errorf("id %d have not found", id)
		}
	}

	return nil
}
