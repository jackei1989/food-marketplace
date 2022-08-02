package main

import (
	"fmt"
	"foodmarketplace/config"
	"foodmarketplace/database"
)

func main() {

	config.SetConfig()
	database.Connect()

	fmt.Println("Server is runnig ...")
}
