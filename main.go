package main

import (
	"ims/database"
	"log"
)

func main() {
	db, err := database.ConnectToDataBase()
	if err != nil {
		log.Fatal("DB connection is Fail!")
	}

	db.Run()
}
