package tests

import (
	"backend-master-class/apis"
	"backend-master-class/db/connection"
	db "backend-master-class/db/sqlc"
	"backend-master-class/util"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

var testQueries *db.Queries
var connectionDB *sql.DB

// func init() {
// 	err := godotenv.Load("../../.env")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

func newTestServer(t *testing.T, store db.Store) *apis.Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := apis.NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	connectionDB = connection.Postgres(config.DBDriver, config.DBSource)

	defer connectionDB.Close()

	gin.SetMode(gin.TestMode)
	testQueries = db.New(connectionDB)
	os.Exit(m.Run())
}
