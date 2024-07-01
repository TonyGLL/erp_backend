package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/TonyGLL/erp_backend/api"
	db "github.com/TonyGLL/erp_backend/db/sqlc"
	"github.com/TonyGLL/erp_backend/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	fmt.Println(config)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store, config.ServerAddress)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
