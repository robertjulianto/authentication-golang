package main

import (
	"fmt"
	"ims/api"
	"ims/database"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	db, err := database.ConnectToDataBase()
	if err != nil {
		log.Fatal("DB connection is Fail!")
	}

	db.Run()

	envFile, _ := godotenv.Read(".env")
	listenAddr := envFile["SERVERADDRESS"]

	server := api.NewServer(listenAddr, db)
	fmt.Println("Server is running on port :", listenAddr)
	log.Fatal(server.Start())
}
