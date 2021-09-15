package main

import (
	"exchange_currency/config"
	"exchange_currency/database"
	"exchange_currency/server"
)

func main() {

	config.Init()
	s := server.NewServer()

	database.Init()

	database.RunMigrations()

	s.Run()

}
