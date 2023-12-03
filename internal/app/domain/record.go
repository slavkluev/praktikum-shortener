package domain

import "context"

// Record хранит данные оригинального URL
type Record struct {
	ID      uint64
	User    string
	URL     string
	Deleted bool
}

// RecordRepository интерфейс репозитория
type RecordRepository interface {
	Ping(ctx context.Context) error
	GetByID(ctx context.Context, id uint64) (Record, error)
	GetByOriginURL(ctx context.Context, originURL string) (Record, error)
	GetByUserID(ctx context.Context, userID string) ([]Record, error)
	Store(ctx context.Context, record *Record) error
	Delete(ctx context.Context, id uint64) error
}

// RecordUsecase интерфейс usecase
type RecordUsecase interface {
	Ping(ctx context.Context) error
	GetByID(ctx context.Context, id uint64) (Record, error)
	GetByUserID(ctx context.Context, id string) ([]Record, error)
	Store(ctx context.Context, record *Record) error
	StoreBatch(ctx context.Context, records []*Record) error
	Delete(ctx context.Context, id uint64) error
	DeleteBatch(ctx context.Context, ids []uint64) error
}
