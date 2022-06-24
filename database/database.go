package database

import (
	"evolve/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	Host = "127.0.0.1"

	DevName     = "evolve"
	DevUsername = "postgres"
	DevPassword = "root"
	DevPort     = "5432"

	ProdDatabaseName = "evolve"
	ProdUsername     = ""
	ProdPassword     = ""
	ProdPort         = "5432"
)

func getDB() string {
	if config.App.ENV == config.Environment.Local {
		return DevName
	}
	return ProdDatabaseName
}
func getUsername() string {
	if config.App.ENV == config.Environment.Local {
		return DevUsername
	}
	return ProdUsername
}
func getPassword() string {
	if config.App.ENV == config.Environment.Local {
		return DevPassword
	}
	return ProdPassword
}

func getPort() string {
	if config.App.ENV == config.Environment.Local {
		return DevPort
	}
	return ProdPort
}

func Start() {

	dsn := "host=" + Host + " user=" + getUsername() + " password=" + getPassword() + " dbname=" + getDB() + " port=" + getPort() + " sslmode=disable TimeZone=Africa/Lagos"

	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}
