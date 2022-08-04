package database

import (
	"fmt"
	"foodmarketplace/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Connection struct {
	Conn *gorm.DB
}

func Connect() Connection {

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v timezone=%v",
		config.AppConfig.DbHost,
		config.AppConfig.DbUser,
		config.AppConfig.DbPsssword,
		config.AppConfig.DbName,
		config.AppConfig.DbProt,
		config.AppConfig.DbTimezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connection ...")

	return Connection{Conn: db}
}
