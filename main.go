package main

import (
	"database/sql"
	"log"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/api"
	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config file", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the database", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
}
