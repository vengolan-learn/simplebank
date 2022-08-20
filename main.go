package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vengolan/simplebank/api"
	db "github.com/vengolan/simplebank/db/sqlc"
	"github.com/vengolan/simplebank/db/util"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configurations")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("Unable to connect to database: \ndbDriver:%s  \ndbSource:%s\n %v", config.DBDriver, config.DBSource, err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
