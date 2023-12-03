package usecase

import (
	"context"
	"errors"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/slavkluev/praktikum-shortener/internal/app/domain"
)

type recordUsecase struct {
	recordRepo     domain.RecordRepository
	contextTimeout time.Duration
}

// NewRecordUsecase создание recordUsecase
func NewRecordUsecase(r domain.RecordRepository, timeout time.Duration) domain.RecordUsecase {
	return &recordUsecase{
		recordRepo:     r,
		contextTimeout: timeout,
	}
}

// Ping используется для проверки доступности сервиса
func (r *recordUsecase) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	return r.recordRepo.Ping(ctx)
}

// GetByID получение записи по ID
func (r *recordUsecase) GetByID(ctx context.Context, id uint64) (domain.Record, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	res, err := r.recordRepo.GetByID(ctx, id)
	if err != nil {
		return domain.Record{}, err
	}

	return res, nil
}

// GetByUserID получение всех записей у пользователя
func (r *recordUsecase) GetByUserID(ctx context.Context, id string) ([]domain.Record, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	res, err := r.recordRepo.GetByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Store вставка одной записи
func (r *recordUsecase) Store(ctx context.Context, record *domain.Record) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	existingRecord, err := r.recordRepo.GetByOriginURL(ctx, record.URL)

	if err == nil {
		record.ID = existingRecord.ID
		return domain.ErrConflict
	}

	if errors.Is(err, domain.ErrNotFound) {
		return r.recordRepo.Store(ctx, record)
	}

	return err
}

// StoreBatch вставка нескольких записей
func (r *recordUsecase) StoreBatch(ctx context.Context, records []*domain.Record) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	grp, ctx := errgroup.WithContext(ctx)

	for _, record := range records {
		record := record
		grp.Go(func() error {
			return r.Store(ctx, record)
		})
	}

	return grp.Wait()
}

// Delete удаление одной записи
func (r *recordUsecase) Delete(ctx context.Context, id uint64) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	return r.recordRepo.Delete(ctx, id)
}

// DeleteBatch удаление нескольких записей
func (r *recordUsecase) DeleteBatch(ctx context.Context, ids []uint64) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	grp, ctx := errgroup.WithContext(ctx)

	for _, id := range ids {
		id := id
		grp.Go(func() error {
			return r.Delete(ctx, id)
		})
	}

	return grp.Wait()
}
