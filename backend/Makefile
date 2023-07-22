.PHONY: clean critic security lint test build run

APP_NAME = oa
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/platform/migrations
DATABASE_URL = postgres://postgres:password@oa-postgres/postgres?sslmode=disable

clean:
	rm -rf ./build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./...

test: clean critic security lint
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: swag build
	$(BUILD_DIR)/$(APP_NAME)

migrate.up:
	docker run --rm -d -v $(MIGRATIONS_FOLDER):/migrations \
		--network dev-network migrate/migrate \
    -path=/migrations/ -database "$(DATABASE_URL)" up

migrate.down:
	docker run --rm -d -v $(MIGRATIONS_FOLDER):/migrations \
		--network dev-network migrate/migrate \
    -path=/migrations/ -database "$(DATABASE_URL)" down

migrate.force:
	docker run --rm -d -v $(MIGRATIONS_FOLDER):/migrations \
		--network dev-network migrate/migrate \
    -path=/migrations/ -database "$(DATABASE_URL)"  force $(version)

docker.run: docker.stop docker.network docker.postgres swag docker.fiber docker.redis migrate.up

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.fiber.build:
	docker build -t fiber .

docker.fiber: docker.fiber.build
	docker run --rm -d \
		--name oa-fiber \
		--network dev-network \
		-p 3000:3000 \
		fiber || :

docker.postgres:
	docker run --rm -d \
		--name oa-postgres \
		--network dev-network \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-e POSTGRES_DB=postgres \
		-v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
		-p 5432:5432 \
		postgres || :

docker.redis:
	docker run --rm -d \
		--name oa-redis \
		--network dev-network \
		-p 6379:6379 \
		redis || :

docker.stop: docker.stop.fiber docker.stop.postgres docker.stop.redis

docker.stop.fiber:
	docker stop oa-fiber || :

docker.stop.postgres:
	docker stop oa-postgres || :

docker.stop.redis:
	docker stop oa-redis || :

swag:
	swag init
