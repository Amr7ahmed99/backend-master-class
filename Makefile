postgres:
	docker run --name posqlDB -p 5430:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it posqlDB createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it posqlDB dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5430/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5430/simple_bank?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown
