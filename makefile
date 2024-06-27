postgres:
	docker run --name postgres-tasktracker -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
migrate:
	soda create -e development
	soda migrate
