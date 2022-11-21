//go:build memdb

package main

import (
	"path/filepath"

	"github.com/nmikal98/cs499-hw3/hotelapp/internal/profile"
)

func initializeProfileDatabase() *profile.DatabaseSession {
	dbsession := profile.NewDatabaseSession()
	dbsession.LoadDataFromJsonFile(filepath.Join(*jsonDataDir, "hotels.json"))
	return dbsession
}
