package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

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

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJson, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJson)
	}
}

func main() {
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe("127.0.0.1:5000", nil)
}
