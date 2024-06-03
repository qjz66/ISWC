package main

import (
	"Server/db"
	"Server/routers"
	"log"
)

func main() {
	db.InitDB()
	db.InitRedis()
	router := routers.InitRouter()
	log.Fatal(router.Run())
}
