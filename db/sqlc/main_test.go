package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var connectionDB *sql.DB

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err.Error())
	}
}

func TestMain(m *testing.M) {

	var err error
	connectionDB, err = sql.Open(os.Getenv("DB_DRIVER_NAME"), os.Getenv("DB_DATA_SOURCE"))
	if err != nil {
		log.Fatal("cannot connect to DB", err)
	}
	testQueries = New(connectionDB)
	os.Exit(m.Run())
}
