package main

import (
	"encoding/json"
	"log"
	"net/http"

	"os"

	"github.com/ctrlaltreboot/igor/ean"
	"github.com/ctrlaltreboot/igor/hotels"
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
		hotelsAPIEndpoint = z
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/ean", ean.List)
	http.HandleFunc("/hotels", hotels.List)
	cheapestHandler := hotels.CheapestHandler{HotelsAPIEndpoint: hotelsAPIEndpoint}

	log.Fatal(http.ListenAndServe("0.0.0.0:8088", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	hello := map[string]string{"hello": "igor"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hello)
}
