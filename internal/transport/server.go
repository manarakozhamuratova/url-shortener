package transport

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
	"urlshortener/internal/config"
	"urlshortener/internal/transport/handlers"
)

type Server struct {
	cfg     *config.Config
	handler *handlers.Handler
	App     *echo.Echo
}

func NewServer(cfg *config.Config, handler *handlers.Handler) *Server {
	return &Server{cfg: cfg, handler: handler}
}

func (s *Server) StartHTTPServer(ctx context.Context) error {
	s.App = s.BuildEngine()

	s.SetupRoutes()

	go func() {
		if err := s.App.Start(fmt.Sprintf("%s", ":9090")); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := s.App.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}
	log.Print("server exited properly")
	return nil
}

func (s *Server) BuildEngine() *echo.Echo {
	e := echo.New()

	return e
}
