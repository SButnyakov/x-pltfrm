MG_PATH=../db/postgres/migrations

ifeq (${OS}, Windows_NT)
	PG_USER=postgres
	PG_PASSWORD=password
	PG_HOST=localhost
	PG_PORT=5432
	PG_NAME=x-pltfrm/music
endif

PG_URL=postgres://${PG_USER}:${PG_PASSWORD}@${PG_HOST}:${PG_PORT}/${PG_NAME}?sslmode=disable

run: build
	./main

build:
	go build -C ./cmd -o ./../main

new_migration:
	migrate create -ext sql -dir ${MG_PATH} -seq ${name}

migrate_up:
	migrate -database ${PG_URL} -path ${MG_PATH} up ${steps}

migrate_down:
	migrate -database ${PG_URL} -path ${MG_PATH} down ${steps}
