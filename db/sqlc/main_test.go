package db

import (
	"backend-master-class/db/connection"
	"backend-master-class/util"
	"database/sql"
	"log"
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
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	connectionDB = connection.Postgres(config.DBDriver, config.DBSource)

	defer connectionDB.Close()

	testQueries = New(connectionDB)
	os.Exit(m.Run())
}
