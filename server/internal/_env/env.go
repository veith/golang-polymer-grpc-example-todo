package environment

import (
	"log"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
)

// Settings, Config, .. können nachher hierrüber allen packages mitgegeben werden
type Environment struct {
	DB sqlbuilder.Database
}

var Env *Environment

func InitEnv() {
	Env = &Environment{}
	dbSession, _ := connectDB()
	Env.DB = dbSession
	// config kann hier nachher an env angehängt werden
}

func connectDB() (sqlbuilder.Database, error) {
	var settings = sqlite.ConnectionURL{
		Database: `data/data.db`, // Path to database file.
	}
	dbSession, err := sqlite.Open(settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}

	return dbSession, err
}
