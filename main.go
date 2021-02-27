package main

import (
	"log"
	"net/http"

	"github.com/hscgavin/inventory-service/product"
	"github.com/hscgavin/inventory-service/receipt"
)

const apiBasePath = "/api"

func main() {
	product.SetupRoutes(apiBasePath)
	receipt.SetupRoutes(apiBasePath)
	log.Fatal(http.ListenAndServe("127.0.0.1:5000", nil))
}
