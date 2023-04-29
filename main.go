package main

import (
	"database/sql"
	"log"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/api"
	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver      = "mysql"
	dbSource      = "root:secret@tcp(localhost:3307)/ismogged?parseTime=true"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	server.Start(serverAddress)
}
