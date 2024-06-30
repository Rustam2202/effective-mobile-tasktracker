db:
	docker rm -f postgres-tasktracker || true
	docker run --name postgres-tasktracker -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
	timeout 2
	make migrate-dev
	make migrate-test

migrate-dev:
	soda create -e development
	soda migrate

migrate-test:
	soda create -e test
	soda migrate --env test

dev:
	buffalo dev
