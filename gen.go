package gen

//go:generate go run ./cmd/tern-dot-env/main.go
//go:generate sqlc generate -f ./internal/pgstore/sqlc.yml
