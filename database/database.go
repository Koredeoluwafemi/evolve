package database

import (
	"evolve/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	DevHost = "127.0.0.1"

	DevName     = "evolve"
	DevUsername = "postgres"
	DevPassword = "root"
	DevPort     = "5432"

	//postgres://autktogywftngg:b7968e3676fc48b20bbda1716b6f65fe2931845690a7420445d1b5be8ca10c80@ec2-99-80-170-190.eu-west-1.compute.amazonaws.com:5432/d1pqvjjggbopo2

	ProdDatabaseName = "d1pqvjjggbopo2"
	ProdUsername     = "autktogywftngg"
	ProdPassword     = "b7968e3676fc48b20bbda1716b6f65fe2931845690a7420445d1b5be8ca10c80"
	ProdPort         = "5432"
	ProdHost         = "ec2-99-80-170-190.eu-west-1.compute.amazonaws.com"

	//mysql://bcd421f86f282d:e8c50bef@eu-cdbr-west-02.cleardb.net/heroku_d9522cccfb35a88?reconnect=true
	//HerokuHost     = "eu-cdbr-west-02.cleardb.net"
	//HerokuName     = "heroku_d9522cccfb35a88"
	//HerokuUsername = "bcd421f86f282d"
	//HerokuPassword = "e8c50bef"
	//HerokuPort     = "3306"
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
func getHost() string {
	if config.App.ENV == config.Environment.Local {
		return DevHost
	}
	return ProdHost
}

func Start() {

	dsn := "host=" + getHost() + " user=" + getUsername() + " password=" + getPassword() + " dbname=" + getDB() + " port=" + getPort() + " sslmode=disable TimeZone=Africa/Lagos"

	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}
