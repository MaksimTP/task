DATABASE_URL=postgres://postgres:postgres@localhost:5432/postgres
EXEC_PATH=cmd/main.go

generate_docs:
	swag init -g $(EXEC_PATH)

run: generate_docs
	go mod tidy && go mod download && \
	go run ./$(EXEC_PATH)

delete_docs:
	rm -rf docs

migrate-up:
	migrate -path migrations/ -database $(DATABASE_URL) up

migrate-down:
	migrate -path migrations/ -database $(DATABASE_URL) down -all