package main

import (
	"net/http"

	"github.com/hscgavin/inventory-service/product"
)

const apiBasePath = "/api"

func main() {
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe("127.0.0.1:5000", nil)
}
