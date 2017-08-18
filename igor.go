package main

import (
	"encoding/json"
	"github.com/ctrlaltreboot/igor/ean"
	"github.com/ctrlaltreboot/igor/hotels"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/ean", ean.List)
	http.HandleFunc("/hotels", hotels.List)
	http.HandleFunc("/cheapest_hotel", hotels.Cheapest)
	log.Fatal(http.ListenAndServe("0.0.0.0:8088", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	hello := map[string]string{"hello": "igor"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hello)
}
