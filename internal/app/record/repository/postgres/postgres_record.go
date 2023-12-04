package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"

	"github.com/slavkluev/praktikum-shortener/internal/app/domain"
)

type postgresRecordRepository struct {
	db *sql.DB
}

// NewPostgresRecordRepository создание postgresRecordRepository
func NewPostgresRecordRepository(db *sql.DB) (domain.RecordRepository, error) {
	p := &postgresRecordRepository{
		db: db,
	}

	err := p.init()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *postgresRecordRepository) init() error {
	_, err := p.db.Exec(`
CREATE TABLE IF NOT EXISTS url
(
    id         bigserial primary key,
    user_id    varchar(36),
    origin_url varchar(255),
    deleted    boolean default false,
    CONSTRAINT origin_url_unique UNIQUE (origin_url)
)
`)

	return err
}

// Ping проверка доступности хранилища
func (p *postgresRecordRepository) Ping(ctx context.Context) error {
	return p.db.PingContext(ctx)
}

// GetByID получение записи по ID
func (p *postgresRecordRepository) GetByID(ctx context.Context, id uint64) (domain.Record, error) {
	var record domain.Record

	row := p.db.QueryRowContext(ctx, "SELECT id, user_id, origin_url, deleted FROM url WHERE id = $1", id)
	err := row.Scan(&record.ID, &record.User, &record.URL, &record.Deleted)
	if err != nil {
		return domain.Record{}, err
	}

	return record, nil
}

// GetByOriginURL получение записи по оригинальному URL
func (p *postgresRecordRepository) GetByOriginURL(ctx context.Context, originURL string) (domain.Record, error) {
	var record domain.Record

	row := p.db.QueryRowContext(ctx, "SELECT id, user_id, origin_url FROM url WHERE origin_url = $1 AND deleted = false", originURL)

	err := row.Scan(&record.ID, &record.User, &record.URL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Record{}, domain.ErrNotFound
		}
		return domain.Record{}, err
	}

	return record, nil
}

// GetByUserID получение всех записей у пользователя
func (p *postgresRecordRepository) GetByUserID(ctx context.Context, userID string) ([]domain.Record, error) {
	var records []domain.Record

	rows, err := p.db.QueryContext(ctx, "SELECT id, user_id, origin_url FROM url WHERE user_id = $1 AND deleted = false", userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var record domain.Record
		err := rows.Scan(&record.ID, &record.User, &record.URL)
		if err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// Store вставка одной записи
func (p *postgresRecordRepository) Store(ctx context.Context, record *domain.Record) error {
	sqlStatement := "INSERT INTO url (user_id, origin_url) VALUES ($1, $2) RETURNING id"

	err := p.db.QueryRowContext(ctx, sqlStatement, record.User, record.URL).Scan(&record.ID)

	if err != nil {
		var pge *pgconn.PgError
		if errors.As(err, &pge) && pge.Code == pgerrcode.UniqueViolation {
			return domain.ErrConflict
		}
		return err
	}

	return nil
}

// Delete удаление одной записи
func (p *postgresRecordRepository) Delete(ctx context.Context, id uint64) error {
	sqlStatement := "UPDATE url SET deleted = true WHERE id = $1"

	_, err := p.db.ExecContext(ctx, sqlStatement, id)

	return err
}

// GetUrlsCount получение количества ссылок
func (p *postgresRecordRepository) GetUrlsCount(ctx context.Context) (uint64, error) {
	sqlStatement := "SELECT COUNT(*) FROM url"

	var urlsCount uint64

	err := p.db.QueryRowContext(ctx, sqlStatement).Scan(&urlsCount)
	if err != nil {
		return 0, err
	}

	return urlsCount, nil
}

// GetUsersCount получение количества пользователей
func (p *postgresRecordRepository) GetUsersCount(ctx context.Context) (uint64, error) {
	sqlStatement := "SELECT COUNT(DISTINCT user_id) FROM url"

	var urlsCount uint64

	err := p.db.QueryRowContext(ctx, sqlStatement).Scan(&urlsCount)
	if err != nil {
		return 0, err
	}

	return urlsCount, nil
}
