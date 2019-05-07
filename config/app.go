package config

import (
	"github.com/spf13/viper"
)

// Version represents app version
var Version = "unversioned"

// Application holds the application configuration
type Application struct {
	Base    string
	Env     string
	Port    int
	Sentry  string
	Users   []User
	Version string
}

// User holds the auth user information
type User struct {
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
}

// app is the default application configuration
var app Application

// App returns the default application configuration
func App() *Application {
	return &app
}

// LoadApp loads application configuration
func LoadApp() {
	mu.Lock()
	defer mu.Unlock()
	env := EnvDevelopment
	if e := viper.GetString("env"); e != "" {
		env = e
	}

	usrs := []User{}
	viper.UnmarshalKey("users", &usrs)
	app = Application{
		Base:    viper.GetString("base"),
		Port:    viper.GetInt("port"),
		Env:     env,
		Sentry:  viper.GetString("sentry.dsn"),
		Users:   usrs,
		Version: Version,
	}
}
