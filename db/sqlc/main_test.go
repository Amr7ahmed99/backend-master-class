package db

import (
	"backend-master-class/db/connection"
	"database/sql"
	"os"
	"testing"
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

	connectionDB = connection.Postgres()

	defer connectionDB.Close()

	testQueries = New(connectionDB)
	os.Exit(m.Run())
}
