//go:build mongodb

package main

import (
	"flag"
	"path/filepath"

	"github.com/nmikal98/cs499-hw3/hotelapp/internal/rate"
)

var (
	database_addr = flag.String("db_addr", "mongodb-rate:27017", "Address of the data base server")
)

func initializeRateDatabase() *rate.DatabaseSession {
	dbsession := rate.NewDatabaseSession(*database_addr)
	dbsession.LoadDataFromJsonFile(filepath.Join(*jsonDataDir, "inventory.json"))
	return dbsession
}
