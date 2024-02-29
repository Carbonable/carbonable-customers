default:
    just --list

db_url := "postgres://carbonable:carbonable@localhost:5432/carbonable_customers?sslmode=disable"

# start docker database
start_db:
    docker compose -f docker-compose.yml up -d

# stop docker database
stop_db:
    docker compose -f docker-compose.yml stop

# starting customers api
api:
  DATABASE_URL={{db_url}} go run cmd/api/main.go

# running migration
migrate:
  DATABASE_URL={{db_url}} go run cmd/migrate/main.go
