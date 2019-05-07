package config

import "github.com/spf13/viper"

// Database holds the database configuration
type Database struct {
	URI string
}

var db Database

// DB returns the default database configuration
func DB() *Database {
	return &db
}

// LoadDB loads database configuration
func LoadDB() {
	mu.Lock()
	defer mu.Unlock()
	db = Database{
		URI: viper.GetString("mongo.uri"),
	}
}
