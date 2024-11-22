package main

import (
	"context"
	"github.com/inchori/starknet-indexer/config"
	"github.com/inchori/starknet-indexer/internal/database"
	"log"
	"net/http"
)

func main() {
	config := config.NewConfig()
	dbClient := database.ConnectDb(config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
	ctx := context.Background()
	if err := dbClient.Schema.Create(ctx); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	if err := http.ListenAndServe(":"+config.Port, nil); err != nil {
		panic(err)
	}
	log.Println("The server is running on " + config.Port + "...")
}
