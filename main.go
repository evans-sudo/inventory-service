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
	log.Fatal(http.ListenAndServe(":5000", nil))
}