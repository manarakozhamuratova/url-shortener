package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	_ "urlshortener/docs"
	"urlshortener/internal/config"
	"urlshortener/internal/repository"
	"urlshortener/internal/service"
	"urlshortener/internal/transport"
	"urlshortener/internal/transport/handlers"
)

// @title Super API
// @version 1.0
// @description This is my first swagger documentation.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:9090

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := repository.NewConnection(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewUrlRepository(db)
	srv := service.New(repo, cfg.ShortURLLen, cfg.ShortUrlDuration)
	handler := handlers.NewHandler(srv)
	server := transport.NewServer(&cfg, handler)
	server.StartHTTPServer(ctx)
}
