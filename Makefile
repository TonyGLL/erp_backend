initschema:
	migrate create -ext sql -dir db/migration -seq init_schema

postgres:
	docker-compose up -d --build db

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
	CONFIG_FILE=local.env go run main.go

build:
	@echo "Building..."

	@CONFIG_FILE=local.env go build -o main main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/TonyGLL/erp_backend/db/sql Store

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    CONFIG_FILE=local.env air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/air-verse/air@latest; \
	        CONFIG_FILE=local.env air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

up-dev:
	@CONFIG_FILE=dev.env docker-compose up --build -d app

up-prod:
	@CONFIG_FILE=prod.env docker-compose up --build -d app

down:
	@docker-compose down

.PHONY: postgres createdb dropdb migrateup migratedown server mock up up-dev up-prod down