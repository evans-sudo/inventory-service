package main

import (
	"inventory/database"
	"inventory/product"
	"inventory/receipt"
	"log"
	"net/http"
)



const basePath = "/api"


func main() {
	database.SetupDatabase()
	receipt.SetupRoutes(basePath)
	product.SetupRoutes(basePath)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}