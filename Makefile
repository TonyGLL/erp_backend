initschema:
	migrate create -ext sql -dir db/migration -seq init_schema

postgres:
	docker run --name postgres_erp -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

createdb:
	docker exec -it postgres_erp createdb --username=root --owner=root erp

dropdb:
	docker exec -it postgres_erp dropdb erp

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/erp?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/erp?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

server:
	go run main.go

build:
	@echo "Building..."

	@go build -o main main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/TonyGLL/erp_backend/db/sqlc Store

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/air-verse/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock