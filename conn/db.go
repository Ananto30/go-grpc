package conn

import (
	mgo "gopkg.in/mgo.v2"
	"github.com/Ananto30/go-grpc/config"
)

// DB holds the database session
type DB struct{ *mgo.Session }

// defaultDB is the default database instance
var defaultDB DB

// Connect sets the db client of database using configuration cfg
func (db *DB) Connect(cfg *config.Database) error {
	// open a database connection using mgo
	s, err := mgo.Dial(cfg.URI)
	if err != nil {
		return err
	}
	db.Session = s
	return nil
}

// DefaultDB returns default db
func DefaultDB() DB {
	return defaultDB
}

// ConnectDB sets the db client of database using default configuration file
func ConnectDB() error {
	return defaultDB.Connect(config.DB())
}
