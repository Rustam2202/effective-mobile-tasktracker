 include .env

db:
# docker rm -f postgres-tasktracker || true
	docker run --name $(CONTAINER_NAME) -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -p $(POSTGRES_PORT):5432 -d postgres
	@$(SLEEP) 2
	make migrate-dev

migrate-dev:
	soda create -e development
	soda migrate --env development
migrate-test:
	soda create -e test
	soda migrate --env test

swag:
	swag fmt
	swag init -o docs -g cmd/app/main.go
	npx @redocly/cli build-docs ./docs/swagger.json -o ./docs/index.html

test:
	buffalo test

dev:
	buffalo dev


ifeq ($(OS),Windows_NT)
    SLEEP = powershell -Command "Start-Sleep -Seconds"
else
    SLEEP = sleep
endif
