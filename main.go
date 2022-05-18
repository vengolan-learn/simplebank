package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vengolan/simplebank/api"
	db "github.com/vengolan/simplebank/db/sqlc"
	"github.com/vengolan/simplebank/db/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("Unable to connect to database: \ndbDriver:%s  \ndbSource:%s\n ", config.DBDriver, config.DBSource)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
