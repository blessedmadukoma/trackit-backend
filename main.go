package main

import (
	"database/sql"
	"fmt"
	"log"
	"trackit/api"
	db "trackit/db/sqlc"
	"trackit/util"

	_ "trackit/docs"

	_ "github.com/lib/pq"
)

//	@title			Trakkit Backend
//	@version		1.0
//	@description	Backend for TrakkIT, a financial management tracking tool
//	@termsOfService	https://trakkit.vercel.app

//	@contact.name	Madukoma Blessed
//	@contact.url	https://mblessed.vercel.app
//	@contact.email	blessedmadukoma@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		https://trackit-blessedmadukoma.koyeb.app
//	@schemes	https
//	@BasePath	/api/
func main() {
	fmt.Println("Hello from TrackIT! Starting Server...")

	// config, err := util.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("Cannot load config:", err)
	// }
	config := util.LoadEnvConfig(".")
	// fmt.Println(config)

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

	// err = server.StartServer(config.ServerAddress)
	err = server.StartServer(config.Port)
	if err != nil {
		log.Fatal("cannot start server!")
	}
}
