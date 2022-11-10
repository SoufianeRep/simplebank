package main

import (
	"database/sql"
	"log"

	"github.com/SoufianeRep/simplebank/api"
	db "github.com/SoufianeRep/simplebank/db/sqlc"
	"github.com/SoufianeRep/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load confi:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Canot connect to DB:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot Start the server:", err)
	}
}
