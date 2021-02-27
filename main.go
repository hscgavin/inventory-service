package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hscgavin/inventory-service/product"
)

var productList []Product

func init() {
	productsJSON := `[
		{
			"productId": 1,
			"manufacturer": "oz",
			"sku": "abc123",
			"upc": "123928104",
			"pricePerUnit": "657.55",
			"quantityOnHand": 1200,
			"productName": "golang"
		},
		{
			"productId": 2,
			"manufacturer": "ha",
			"sku": "abc888",
			"upc": "1239029348",
			"pricePerUnit": "67.55",
			"quantityOnHand": 8200,
			"productName": "javascript"
		}
	]`
	err := json.Unmarshal([]byte(productsJSON), &productList)
	if err != nil {
		log.Fatal(err)
	}
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("called"))
}

func getNextID() int {
	highestID := -1
	for _, prod := range productList {
		if highestID < prod.ProductID {
			highestID = prod.ProductID
		}
	}
	return highestID + 1
}

func findProductByID(productID int) (*Product, int) {
	for i, product := range productList {
		if product.ProductID == productID {
			return &product, i
		}
	}
	return nil, 0
}

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before handler; middleware start")
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("middleware finished; %s", time.Since(start))
	})
}

const apiBasePath = "/api"

func main() {
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe("127.0.0.1:5000", nil)
}
