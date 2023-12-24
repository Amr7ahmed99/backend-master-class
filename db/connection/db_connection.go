package connection

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Postgres(driver string, source string) *sql.DB {

	connectionDB, err := sql.Open(driver, source)
	if err != nil {
		log.Fatal("cannot connect to DB", err)
	}

	return connectionDB
}
