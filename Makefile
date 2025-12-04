# Makefile for golang-migrate (PostgreSQL / Neon)

# Allows creating migrations, running up/down, forcing, and resetting.

MIGRATION_PATH := 
DB_URL := postgresql://neondb_owner:npg_k0WD6GpeTSui@ep-dawn-leaf-ad26dewr-pooler.c-2.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require

.PHONY: migrate-create migrate-up migrate-down migrate-force migrate-reset

# Create a new migration

# Usage: make migrate-create NAME=create_users_table

migrate-create:
@migrate create -seq -ext sql -dir $(MIGRATION_PATH) $(NAME)

# Apply all pending migrations

migrate-up:
@migrate -path=$(MIGRATION_PATH) -database='$(DB_URL)' up

# Rollback last migration (down)

migrate-down:
@migrate -path=$(MIGRATION_PATH) -database='$(DB_URL)' down 1

# Force migration version (use carefully)

# Usage: make migrate-force VERSION=0

migrate-force:
@migrate -path=$(MIGRATION_PATH) -database='$(DB_URL)' force $(VERSION)

# Reset all migrations: force version 0 and reapply everything

migrate-reset:
@migrate -path=$(MIGRATION_PATH) -database='$(DB_URL)' force 0
@migrate -path=$(MIGRATION_PATH) -database='$(DB_URL)' up

.PHONY: seed
seed:
@go run cmd/migrate/seed/main.go

.PHONY: gem-docs
gen-docs:
@swag init -g ./api/main.go -d cmd,internal && swag fmt
