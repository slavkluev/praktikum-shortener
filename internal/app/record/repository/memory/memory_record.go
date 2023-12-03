package memory

import (
	"context"
	"sync"

	"github.com/slavkluev/praktikum-shortener/internal/app/domain"
)

type memoryRecordRepository struct {
	Start   uint64
	Records map[uint64]domain.Record
	mu      sync.Mutex
}

// NewMemoryRecordRepository создание memoryRecordRepository
func NewMemoryRecordRepository() domain.RecordRepository {
	return &memoryRecordRepository{
		Start:   1,
		Records: make(map[uint64]domain.Record),
	}
}

// Ping проверка доступности хранилища
func (m *memoryRecordRepository) Ping(ctx context.Context) error {
	return nil
}

// GetByID получение записи по ID
func (m *memoryRecordRepository) GetByID(ctx context.Context, id uint64) (domain.Record, error) {
	if record, ok := m.Records[id]; ok {
		return record, nil
	}

	return domain.Record{}, domain.ErrNotFound
}

// GetByOriginURL получение записи по оригинальному URL
func (m *memoryRecordRepository) GetByOriginURL(ctx context.Context, originURL string) (domain.Record, error) {
	for _, record := range m.Records {
		if record.URL == originURL && !record.Deleted {
			return record, nil
		}
	}

	return domain.Record{}, domain.ErrNotFound
}

// GetByUserID получение всех записей у пользователя
func (m *memoryRecordRepository) GetByUserID(ctx context.Context, userID string) ([]domain.Record, error) {
	var records []domain.Record

	for _, record := range m.Records {
		if record.User == userID && !record.Deleted {
			records = append(records, record)
		}
	}

	return records, nil
}

// Store вставка одной записи
func (m *memoryRecordRepository) Store(ctx context.Context, record *domain.Record) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	record.ID = m.Start
	m.Start++

	m.Records[record.ID] = *record

	return nil
}

// Delete удаление одной записи
func (m *memoryRecordRepository) Delete(ctx context.Context, id uint64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	record, ok := m.Records[id]
	if !ok {
		return domain.ErrNotFound
	}

	record.Deleted = true
	m.Records[record.ID] = record

	return nil
}
