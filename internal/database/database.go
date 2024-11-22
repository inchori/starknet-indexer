package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/inchori/starknet-indexer/internal/ent"
	"log"
)

var DBConn *ent.Client

func ConnectDb(dbUser, dbPassword, dbHost, dbPort, dbName string) *ent.Client {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", dbUser, dbPassword, dbHost, dbPort, dbName)
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	DBConn = client
	return DBConn
}
