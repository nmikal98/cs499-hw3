//go:build memdb

package main

import (
	"path/filepath"

	"github.com/nmikal98/cs499-hw3/hotelapp/internal/geo"
)

func initializeGeoDatabase() *geo.DatabaseSession {
	dbsession := geo.NewDatabaseSession()
	dbsession.LoadDataFromJsonFile(filepath.Join(*jsonDataDir, "geo.json"))
	return dbsession
}
