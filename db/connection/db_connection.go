package connection

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Postgres() *sql.DB {
	connectionDB, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to DB", err)
	}

	return connectionDB
}
