postgres:
	docker run --name postgres12 --network bank-network -p 5430:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

localpsql_createdb:
	PGPASSWORD=123456789 psql -U postgres -p 5432 -c "create database simple_bank;"

localpsql_dropdb:
	PGPASSWORD=123456789 psql -U postgres -p 5432 -c "drop database simple_bank;"

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

localpsql_migrateup:
	migrate -path db/migration -database "postgresql://postgres:123456789@localhost:5432/simple_bank?sslmode=disable" -verbose up

localpsql_migratedown:
	migrate -path db/migration -database "postgresql://postgres:123456789@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5430/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5430/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5430/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5430/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./tests/...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go backend-master-class/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock

.LOCAL_PHONY: localpsql_createdb localpsql_dropdb localpsql_migrateup localpsql_migratedown