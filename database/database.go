package database

import (
	"fmt"
	"foodmarketplace/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	Conn *gorm.DB
}

func Connect() Connection {

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v timezone=%v",
		config.AppConfig.DB_HOST,
		config.AppConfig.DB_USER,
		config.AppConfig.DB_PSSSWORD,
		config.AppConfig.DB_NAME,
		config.AppConfig.DB_PROT,
		config.AppConfig.DB_TIMEZONE,
	)

	// fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connection ...")

	return Connection{Conn: db}
}
