GOLANGCI_LINT = $(HOME)/bin/golangci-lint

.PHONY: lint
lint:
	### RUN GOLANGCI-LINT ###
	$(GOLANGCI_LINT) run ./... --config=golangci.yaml

.PHONY: lint-fast
lint-fast:
	$(GOLANGCI_LINT) run ./... --fast --config=golangci.yaml

migration_create:
	migrate create -ext sql -dir db/migrations/ -seq $(NAME)

migration_version:
	migrate -path db/migrations/ -database "postgresql://postgres:postgres@localhost:5432/postgres_db?sslmode=disable" -verbose version

migration_up:
	migrate -path db/migrations/ -database "postgresql://postgres:postgres@localhost:5432/postgres_db?sslmode=disable" up

migration_down:
	migrate -path db/migrations/ -database "postgresql://postgres:postgres@localhost:5432/postgres_db?sslmode=disable" -verbose down

migration_fix:
	migrate -path db/migrations/ -database "postgresql://postgres:postgres@localhost:5432/postgres_db?sslmode=disable" force $(VERSION)