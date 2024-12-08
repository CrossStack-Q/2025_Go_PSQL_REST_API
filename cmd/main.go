package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/CrossStack-Q/2025_Go_PSQL_REST_API.git/cmd/api"
	"github.com/CrossStack-Q/2025_Go_PSQL_REST_API.git/config"
	"github.com/CrossStack-Q/2025_Go_PSQL_REST_API.git/db"
)

func initDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}

	log.Println("Connected to DB ", config.Envs.DBName)

	return nil
}

func main() {
	fmt.Println("Hello Ji")
	// initialize DB
	dbConn, err := db.NewPostgreSQL(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Envs.DBHost,
		config.Envs.DBPort,
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBName,
	))
	if err != nil {
		log.Println("error in postgres conn", err)
	}

	if err := initDB(dbConn); err != nil {
		log.Println(err)
		return
	}
	// start api server
	apiServer := api.NewAPISERVER(":8080")
	if err := apiServer.Run(); err != nil {
		log.Print(err)
	}

}
