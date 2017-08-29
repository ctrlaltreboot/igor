package cheapest

import (
	"encoding/json"
	"fmt"
	"igor/ean"
	"igor/hotels"
	"net/http"
	"sort"

	"igor/helper"
)

type Cheapest struct {
	PropertyOffer PropertyOffer `json:"PropertyOffer"`
}

type PropertyOffer struct {
	PropertyName  string  `json:"property_name"`
	RoomTypeName  string  `json:"room_type_name"`
	OfferName     string  `json:"offer_name"`
	Amount        float64 `json:"amount"`
	NonRefundable bool    `json:"is_non_refundable"`
}

type CheapestHandler struct {
	HotelsAPIEndpoint string
	EanAPIEndpoint    string
}

func (u *CheapestHandler) fetchCheapest(out chan<- PropertyOffer) {
	out <- PropertyOffer(hotels.GetCheapest(u.HotelsAPIEndpoint))
	out <- PropertyOffer(ean.GetCheapest(u.EanAPIEndpoint))
	close(out)
}

func cheapest(in <-chan PropertyOffer, out chan<- Cheapest) {
	for v := range in {
		var offer []PropertyOffer
		offer = append(offer, v)
		sort.SliceStable(offer, func(i, j int) bool {
			return offer[i].Amount < offer[j].Amount
		})
		out <- Cheapest{offer[0]}
	}

	close(out)
}

func (u *CheapestHandler) Hotel(w http.ResponseWriter, r *http.Request) {
	res := Cheapest{PropertyOffer(hotels.GetCheapest(u.HotelsAPIEndpoint))}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "error writing JSON response", http.StatusInternalServerError)
		return
	}
}

func (u *CheapestHandler) Ean(w http.ResponseWriter, r *http.Request) {
	res := Cheapest{PropertyOffer(ean.GetCheapest(u.EanAPIEndpoint))}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "error writing JSON response", http.StatusInternalServerError)
		return
	}
}

func (u *CheapestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	offers := make(chan PropertyOffer)
	dCheapest := make(chan Cheapest)

	go u.fetchCheapest(offers)
	go cheapest(offers, dCheapest)

	res := <-dCheapest

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "error writing JSON response", http.StatusInternalServerError)
		return
	}
}

func (u *CheapestHandler) ListHotel(w http.ResponseWriter, r *http.Request) {
	b, err := helper.Fetch(u.HotelsAPIEndpoint)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", b)
}

func (u *CheapestHandler) ListEan(w http.ResponseWriter, r *http.Request) {
	b, err := helper.Fetch(u.EanAPIEndpoint)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", b)
}
