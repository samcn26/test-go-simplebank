package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/samcn26/test-go-simplebank/api"
	db "github.com/samcn26/test-go-simplebank/db/sqlc"
	"github.com/samcn26/test-go-simplebank/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
