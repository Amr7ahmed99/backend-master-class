package main

import (
	"backend-master-class/api"
	"backend-master-class/db/connection"
	db "backend-master-class/db/sqlc"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	var connectionDB *sql.DB
	const serverAddress = "0.0.0.0:8080"
	connectionDB = connection.Postgres()
	store := db.NewStore(connectionDB)
	server := api.NewServer(store)
	if err := server.Start(serverAddress); err != nil {
		log.Fatal("can not start server:", err)
	}
}
