package config

import(
	"log"
	"os"
	"github.com/joho/godotenv"
)

type config struct {
	Port			string `json:"port"`
	DB_HOST			string `json:"db_host"`
	DB_PROT			string `json:"db_prot"`
	DB_PSSSWORD		string `json:"db_password"`
	DB_USER			string `json:"db_user"`
	DB_NAME			string `json:"db_name"`
	DB_TIMEZONE		string `json:"db_timezone"`
}

var AppConfig config

func SetConfig() {
	
	err := godotenv.Load("./config/.env")
	if err!= nil{
		log.Fatal(err)
	}

	AppConfig.Port = os.Getenv("Port")
	AppConfig.DB_HOST = os.Getenv("DB_HOST")
	AppConfig.DB_PROT = os.Getenv("DB_PORT")
	AppConfig.DB_PSSSWORD = os.Getenv("DB_PASSWORD")
	AppConfig.DB_NAME = os.Getenv("DB_NAME")
	AppConfig.DB_TIMEZONE = os.Getenv("DB_TIMEZONE")
}