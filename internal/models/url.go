package models

import (
	"net/url"
	"regexp"
	"strings"
	"time"
)

type URL struct {
	ID        int64      `json:"id" db:"id"`
	IsDeleted bool       `json:"is_deleted" db:"is_deleted"`
	ShortURL  string     `json:"short_url" db:"short_url"`
	LongURL   string     `json:"long_url" db:"long_url"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
	ExpiresAt time.Time  `json:"expires_at" db:"expires_at"`
}

type InsertURL struct {
	ShortURL  string    `json:"short_url"`
	LongURL   string    `json:"long_url"`
	ExpiresAt time.Time `json:"expires_at"`
}

func ValidateURL(u string) bool {
	u = strings.TrimSpace(u)

	parsedURL, err := url.ParseRequestURI(u)
	if err != nil {
		return false
	}

	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}

	re := regexp.MustCompile(`^https?://`)
	if !re.MatchString(u) {
		return false
	}

	return true
}
