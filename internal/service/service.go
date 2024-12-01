package service

import (
	"context"
	"errors"
	"fmt"
	"time"
	"urlshortener/internal/models"
	"urlshortener/internal/random"
	"urlshortener/internal/repository"
)

var (
	ErrLinkNotFound = errors.New("url not found")
	ErrStatNotFound = errors.New("stat not found")
)

type Service struct {
	repo             Repository
	shortURLLen      int
	shortURLDuration time.Duration
}

type Repository interface {
	Create(ctx context.Context, url models.InsertURL) error
	GetURLs(ctx context.Context) ([]models.URL, error)
	GetURL(ctx context.Context, shortUrl string) (*models.URL, error)
	DeleteURL(ctx context.Context, shortURL string) error
	UpsertStats(ctx context.Context, stat models.UpsertStat) error
	GetStats(ctx context.Context, urlID int64) (*models.Stat, error)
}

func New(repo Repository, shortURLLen int, shortURLDuration time.Duration) *Service {
	return &Service{repo: repo, shortURLLen: shortURLLen, shortURLDuration: shortURLDuration}
}

// - Пользователь вводит длинный URL в формате JSON и система генерирует короткий уникальный идентификатор для ссылки в ответ.
// - Важно защитить систему от создания дублирующихся коротких ссылок.
// - Реализовать ограничение по времени жизни коротких ссылок (например, 30 дней).
func (s *Service) Create(ctx context.Context, longURL string) (string, error) {
	shortURL := random.NewRandomString(s.shortURLLen)

	return shortURL, s.repo.Create(ctx, models.InsertURL{
		ShortURL:  shortURL,
		LongURL:   longURL,
		ExpiresAt: time.Now().Add(s.shortURLDuration),
	})
}

// - Пользователь должен иметь возможность просматривать список своих созданных ссылок.
func (s *Service) GetURLs(ctx context.Context) ([]models.URL, error) {
	return s.repo.GetURLs(ctx)
}

// - При обращении к короткой ссылке пользователь должен быть перенаправлен на соответствующий длинный URL.
func (s *Service) GetLongURL(ctx context.Context, shortURL string) (string, error) {
	url, err := s.repo.GetURL(ctx, shortURL)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return "", ErrLinkNotFound
		}

		return "", fmt.Errorf("get url: %w", err)
	}

	if err := s.repo.UpsertStats(ctx, models.UpsertStat{
		URLID:     url.ID,
		VisitedAt: time.Now(),
	}); err != nil {
		return "", err
	}

	return url.LongURL, nil
}

// - Пользователю должен быть доступен функционал для удаления своей короткой ссылки.
func (s *Service) DeleteURL(ctx context.Context, url string) error {
	if err := s.repo.DeleteURL(ctx, url); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrLinkNotFound
		}

		return err
	}

	return nil
}

// - Пользователю должна быть доступна статистика по учету коль-во переходов по своим коротким ссылкам, дата последнего перехода.
func (s *Service) GetStats(ctx context.Context, shortURL string) (*models.Stat, error) {
	url, err := s.repo.GetURL(ctx, shortURL)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrLinkNotFound
		}

		return nil, err
	}

	stat, err := s.repo.GetStats(ctx, url.ID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrStatNotFound
		}

		return nil, err
	}

	return stat, nil
}
