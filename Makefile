.PHONY: run
run:
	go run cmd/marketplace/main.go

.PHONY: build
build:
	go build -o ./bin/app ./cmd/marketplace/main.go

.PHONY: build-app-image
build-app-image:
	docker build -t marketplace:$(version) .

.PHONY: compose-up
compose-up: build-app-image
	docker compose up -d

.PHONY: create-migration
create-migration:
	tern new -m migrations/ $(name)

.PHONY: install-dotenv
install-dotenv:
	npm install -g dotenv-cli

.PHONY: install-tern
install-tern:
	go install github.com/jackc/tern/v2@latest

.PHONY: migrate
migrate:
	dotenv -- tern migrate -m migrations/

.PHONY: rollback
rollback:
	dotenv -- tern migrate -m migrations/ -d -1

.PHONY: test
test:
	go test -coverpkg=./... -coverprofile=c.out.tmp ./...

.PHONY: cover
cover: test
	go tool cover -func=c.out.tmp
