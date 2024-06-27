db:
	docker run --name postgres-tasktracker -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
	sleep 1
	make migrate

migrate:
	soda create -e development
	soda create -e test
	soda migrate

dev:
	buffalo dev
