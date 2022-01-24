package storages

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type FileStorage struct {
	Start uint64
	File  *os.File
}

type Record struct {
	ID  uint64
	URL string
}

func CreateFileStorage(filename string) (*FileStorage, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	if err != nil {
		return nil, err
	}

	var lastID uint64 = 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Bytes()

		record := Record{}
		err := json.Unmarshal(data, &record)
		if err != nil {
			return nil, err
		}

		if record.ID > lastID {
			lastID = record.ID
		}
	}

	return &FileStorage{
		Start: lastID + 1,
		File:  file,
	}, nil
}

func (f *FileStorage) Get(id uint64) (string, error) {
	_, err := f.File.Seek(0, io.SeekStart)

	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(f.File)

	for scanner.Scan() {
		data := scanner.Bytes()

		record := Record{}
		err := json.Unmarshal(data, &record)
		if err != nil {
			return "", err
		}

		if record.ID == id {
			return record.URL, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("id %d have not found", id)
}

func (f *FileStorage) Put(url string) (uint64, error) {
	id := f.Start
	f.Start++

	var record = Record{
		ID:  id,
		URL: url,
	}
	data, err := json.Marshal(&record)

	if err != nil {
		return 0, err
	}

	data = append(data, '\n')
	_, err = f.File.Write(data)

	return id, err
}

func (f *FileStorage) Close() error {
	return f.File.Close()
}
