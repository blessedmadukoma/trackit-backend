package main

import (
	"database/sql"
	"fmt"
	"log"
	"trackit/api"
	db "trackit/db/sqlc"
	"trackit/util"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello from TrackIT! Starting Server...")

	// config, err := util.LoadConfig(".")
	config, err := util.LoadEnvConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	// connect to database
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.StartServer(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server!")
	}
}
