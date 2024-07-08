package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/TonyGLL/erp_backend/api"
	db "github.com/TonyGLL/erp_backend/db/sql"
	"github.com/TonyGLL/erp_backend/util"
	_ "github.com/lib/pq"
)

// @contact.name Tony Gonzalez
// @contact.utl https://github.com/TonyGLL
// @contact.email tonygllambia@gmail.com
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
func main() {
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		log.Fatal("CONFIG_FILE environment variable not set")
	}

	config, err := util.LoadConfig(".", configFile)
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store, config.ServerAddress, config.Version)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
