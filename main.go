package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/HemanthKumarKola1/simple-bank/api"
	db "github.com/HemanthKumarKola1/simple-bank/db/sqlc"
)

const (
	dbDriver     = "postgres"
	dbSource     = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serveAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serveAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
