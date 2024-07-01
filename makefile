db:
	docker rm -f postgres-tasktracker || true
	docker run --name postgres-tasktracker -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
	sleep 2
	make migrate-dev
	make migrate-test

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

dev:
	buffalo dev
