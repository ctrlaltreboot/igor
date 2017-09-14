package main

import (
	"encoding/json"
	"log"
	"net/http"

	"os"

	"igor/cheapest"
)

var (
	hotelsAPIEndpoint = "http://127.0.0.1:5091/hotels"
	eanAPIEndpoint    = "http://127.0.0.1:5092/ean"
)

func init() {
	if v, ok := os.LookupEnv("HOTELS_API_ENDPOINT"); ok {
		hotelsAPIEndpoint = v
	}

	if z, ok := os.LookupEnv("EAN_API_ENDPOINT"); ok {
		eanAPIEndpoint = z
	}
}

func main() {
	http.HandleFunc("/hello", hello)

	cheapestHandler := cheapest.CheapestHandler{
		HotelsAPIEndpoint: hotelsAPIEndpoint,
		EanAPIEndpoint:    eanAPIEndpoint,
	}

	http.HandleFunc("/cheapest_hotel", cheapestHandler.Hotel)
	http.HandleFunc("/hotels", cheapestHandler.ListHotel)

	http.HandleFunc("/cheapest_ean", cheapestHandler.Ean)
	http.HandleFunc("/ean", cheapestHandler.ListEan)

	http.HandleFunc("/cheapest", cheapestHandler.ServeHTTP)

	log.Fatal(http.ListenAndServe("0.0.0.0:5000", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	hello := map[string]string{"hello": "igor"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hello)
}
