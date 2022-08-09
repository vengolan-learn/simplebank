package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vengolan/simplebank/api"
	db "github.com/vengolan/simplebank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("Unable to connect to database: \ndbDriver:%s  \ndbSource:%s\n ", dbDriver, dbSource)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
