migrate:
	migrate -path ./schema -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' up
postgres:
	docker run --name=people_api -e POSTGRES_PASSWORD='postgres' -p 5432:5432 -d --rm postgres
