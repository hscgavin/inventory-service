package main

import (
	"net/http"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("called"))

}
func main() {
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe("127.0.0.1:5000", nil)
}
