//go:build memdb

package main

import (
	"path/filepath"

	"github.com/nmikal98/cs499-hw3/hotelapp/internal/rate"
)

func initializeRateDatabase() *rate.DatabaseSession {
	dbsession := rate.NewDatabaseSession()
	dbsession.LoadDataFromJsonFile(filepath.Join(*jsonDataDir, "inventory.json"))
	return dbsession
}
