package main

import (
	"Server/db"
	"Server/routers"
	"log"
)

func main() {
	db.InitDB()
	router := routers.InitRouter()
	log.Fatal(router.Run())
}
