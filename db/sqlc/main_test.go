package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/SoufianeRep/simplebank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Test :: Couldnt load env config:", err)
	}

	// Establishes a connection to the database
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Canot connect to DB:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
