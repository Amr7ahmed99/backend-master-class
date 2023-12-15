package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var connectionDB *sql.DB

// func init() {
// 	err := godotenv.Load("../../.env")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

func TestMain(m *testing.M) {

	var err error
	connectionDB, err = sql.Open("postgres", "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to DB", err)
	}

	defer connectionDB.Close()

	testQueries = New(connectionDB)
	os.Exit(m.Run())
}
