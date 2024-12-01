# URL Shortener Service

A simple and efficient URL shortening service written in Go. This service generates short URLs for given long URLs, stores them, and provides redirection to the original URLs.

## **How to start?**

```
docker compose up -d 
export POSTGRESQL_URL='postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable'
migrate -database ${POSTGRESQL_URL} -path internal/repository/migrations up
go run cmd/main.go
go test ./tests
```