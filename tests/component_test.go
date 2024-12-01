package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
	"urlshortener/internal/config"
	"urlshortener/internal/models"
	"urlshortener/internal/repository"
	"urlshortener/internal/service"
	"urlshortener/internal/transport"
	"urlshortener/internal/transport/handlers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComponent(t *testing.T) {
	ctx := context.Background()

	cfg := config.Config{
		DBHost:           "localhost",
		DBPort:           "5432",
		DBUser:           "postgres",
		DBPassword:       "postgres",
		DBName:           "postgres",
		SSLMode:          "disable",
		ShortURLLen:      6,
		ShortUrlDuration: time.Hour * 24 * 30,
	}

	db, err := repository.NewConnection(&cfg)
	if err != nil {
		assert.NoError(t, err)
	}
	defer db.Close()

	repo := repository.NewUrlRepository(db)
	srv := service.New(repo, cfg.ShortURLLen, cfg.ShortUrlDuration)
	handler := handlers.NewHandler(srv)
	server := transport.NewServer(&cfg, handler)

	go server.StartHTTPServer(ctx)

	waitForHttpServer(t)

	// create a short url
	shortURL := assertShortURLCreated(t)
	// get list of short urls
	assertShortURLExist(t, shortURL)
	// redirect works correcly
	assertRedirect(t, shortURL)
	assertRedirect(t, shortURL)
	// get stats
	assertStats(t, shortURL)
	// delete short url
	assertShortURLDeleted(t, shortURL)
}

func waitForHttpServer(t *testing.T) {
	t.Helper()

	require.EventuallyWithT(
		t,
		func(t *assert.CollectT) {
			resp, err := http.Get("http://localhost:9090/health")
			if !assert.NoError(t, err) {
				return
			}
			defer resp.Body.Close()

			if assert.Less(t, resp.StatusCode, 300, "API not ready, http status: %d", resp.StatusCode) {
				return
			}
		},
		time.Second*10,
		time.Millisecond*50,
	)
}

func assertShortURLCreated(t *testing.T) string {
	t.Helper()

	req := handlers.CreateShortUrlRequest{
		URL: "https://www.google.com",
	}
	reqBody, err := json.Marshal(req)
	require.NoError(t, err)
	resp, err := http.Post("http://localhost:9090/shortener", "application/json", bytes.NewReader(reqBody))
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	shortURLResp := handlers.CreateShortUrlResponse{}
	err = json.NewDecoder(resp.Body).Decode(&shortURLResp)
	require.NoError(t, err)

	return shortURLResp.ShortUrl
}

func assertShortURLExist(t *testing.T, shortURL string) {
	t.Helper()

	resp, err := http.Get("http://localhost:9090/shortener")
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	shortURLs := []models.URL{}
	err = json.NewDecoder(resp.Body).Decode(&shortURLs)
	require.NoError(t, err)
	require.NotEmpty(t, shortURLs)

	var found bool
	for _, u := range shortURLs {
		if u.ShortURL == shortURL {
			found = true
			break
		}
	}
	require.True(t, found)
}

func assertRedirect(t *testing.T, shortURL string) {
	t.Helper()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Get("http://localhost:9090/" + shortURL)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusFound, resp.StatusCode)
}

func assertStats(t *testing.T, shortURL string) {
	t.Helper()

	resp, err := http.Get("http://localhost:9090/stats/" + shortURL)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	stat := models.Stat{}
	err = json.NewDecoder(resp.Body).Decode(&stat)
	require.NoError(t, err)
	require.NotEmpty(t, stat)

	require.Equal(t, stat.VisitedCount, 2)
}

func assertShortURLDeleted(t *testing.T, shortURL string) {
	t.Helper()

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:9090/%s", shortURL), nil)
	require.NoError(t, err)
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	resp, err = http.Get("http://localhost:9090/shortener")
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	shortURLs := []models.URL{}
	err = json.NewDecoder(resp.Body).Decode(&shortURLs)
	require.NoError(t, err)
	if len(shortURLs) == 0 {
		return
	}
	var found bool
	for _, u := range shortURLs {
		if u.ShortURL == shortURL {
			found = true
			break
		}
	}
	require.False(t, found)
}
