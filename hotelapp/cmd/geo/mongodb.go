//go:build mongodb

package main

import (
	"flag"
	"path/filepath"

	"github.com/nmikal98/cs499-hw3/hotelapp/internal/geo"
)

var (
	database_addr = flag.String("db_addr", "mongodb-geo:27017", "Address of the data base server")
)

func initializeGeoDatabase() *geo.DatabaseSession {
	dbsession := geo.NewDatabaseSession(*database_addr)
	dbsession.LoadDataFromJsonFile(filepath.Join(*jsonDataDir, "geo.json"))
	return dbsession
}
