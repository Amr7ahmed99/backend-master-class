package main

import (
	"backend-master-class/api"
	"backend-master-class/db/connection"
	db "backend-master-class/db/sqlc"
	"backend-master-class/util"
	"database/sql"
	"log"

	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func main() {
	var connectionDB *sql.DB
	config, err := util.LoadConfig(".")
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found:", err)
		} else {
			log.Fatal("Config file was found but another error was produced:", err)
		}
	}
	connectionDB = connection.Postgres(config.DBDriver, config.DBSource)
	store := db.NewStore(connectionDB)
	server := api.NewServer(store)
	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("can not start server:", err)
	}
}
