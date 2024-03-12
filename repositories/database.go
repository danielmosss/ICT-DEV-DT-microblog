package repositories

import (
	"github.com/che-ict/DEV-DT-Microblog/models"
	"github.com/glebarez/sqlite"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var connection *gorm.DB

func connect() {
	var err error
	if os.Getenv("APP_MODE") == "test" {
		connection, err = gorm.Open(mysql.Open(os.Getenv("DB_CONNECTION_STRING")))
	} else {
		connection, err = gorm.Open(sqlite.Open("db.sqlite"))
	}

	if err != nil {
		log.Panic(err)
	}

	err = connection.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		log.Panic(err)
	}
}

func DB() *gorm.DB {
	if connection == nil {
		connect()
	}
	return connection
}
