package repository

import (
	"context"
	"database/sql"
	"errors"
	"urlshortener/internal/models"

	"github.com/jmoiron/sqlx"
)

var (
	ErrNotFound = errors.New("url not found")
)

type UrlRepository struct {
	db *sqlx.DB
}

func NewUrlRepository(db *sqlx.DB) *UrlRepository {
	return &UrlRepository{db: db}
}

func (u *UrlRepository) Create(ctx context.Context, url models.InsertURL) error {
	const query = `
		insert into short_url (long_url, short_url, expires_at) values ($1, $2, $3)
		on conflict do nothing
	`

	if _, err := u.db.ExecContext(ctx, query, url.LongURL, url.ShortURL, url.ExpiresAt); err != nil {
		return err
	}

	return nil
}

func (u *UrlRepository) GetURLs(ctx context.Context) ([]models.URL, error) {
	var url []models.URL
	err := u.db.SelectContext(ctx, &url, "select * from short_url where is_deleted = false and expires_at > now() order by created_at desc")
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (u *UrlRepository) GetURL(ctx context.Context, shortUrl string) (*models.URL, error) {
	var longURL models.URL

	const query = `
		select * from short_url 
		where short_url = $1 and is_deleted = false and expires_at > now()
	`

	err := u.db.GetContext(ctx, &longURL, query, shortUrl)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &longURL, nil
}

func (u *UrlRepository) DeleteURL(ctx context.Context, shortURL string) error {
	const query = `
		update short_url set is_deleted = true, deleted_at = now() where short_url = $1
	`
	if _, err := u.db.ExecContext(ctx, query, shortURL); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFound
		}

		return err
	}

	return nil
}

func (u *UrlRepository) UpsertStats(ctx context.Context, stat models.UpsertStat) error {
	const query = `
		insert into stats (url_id, visited_count, visited_at) values ($1, $2, $3)
		on conflict (url_id) do update set visited_count = stats.visited_count + $2, visited_at = $3
	`

	if _, err := u.db.ExecContext(ctx, query, stat.URLID, 1, stat.VisitedAt); err != nil {
		return err
	}

	return nil
}

func (u *UrlRepository) GetStats(ctx context.Context, urlID int64) (*models.Stat, error) {
	var stat models.Stat
	err := u.db.GetContext(ctx, &stat, "select * from stats where url_id = $1", urlID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &stat, nil
}
