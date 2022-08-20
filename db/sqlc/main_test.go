package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/vengolan/simplebank/db/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("unable to load config ", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("Unable to connect to database: \ndbDriver:%s  \ndbSource:%s\n ", config.DBDriver, config.DBSource)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
