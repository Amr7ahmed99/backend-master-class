package connection

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var ConnectionDB *sql.DB

func Postgres(driver string, source string) *sql.DB {
	if ConnectionDB == nil {
		var err error
		ConnectionDB, err = sql.Open(driver, source)
		if err != nil {
			log.Fatal("cannot connect to DB", err)
		}
	}
	return ConnectionDB
}
