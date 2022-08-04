package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type config struct {
	Port       string `json:"port"`
	DbHost     string `json:"db_host"`
	DbProt     string `json:"db_prot"`
	DbPsssword string `json:"db_password"`
	DbUser     string `json:"db_user"`
	DbName     string `json:"db_name"`
	DbTimezone string `json:"db_timezone"`
}

var AppConfig config

func SetConfig() {

	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal(err)
	}

	AppConfig.Port = os.Getenv("PORT")
	AppConfig.DbHost = os.Getenv("DB_HOST")
	AppConfig.DbUser = os.Getenv("DB_USER")
	AppConfig.DbProt = os.Getenv("DB_PORT")
	AppConfig.DbPsssword = os.Getenv("DB_PASSWORD")
	AppConfig.DbName = os.Getenv("DB_NAME")
	AppConfig.DbTimezone = os.Getenv("DB_TIMEZONE")
}
