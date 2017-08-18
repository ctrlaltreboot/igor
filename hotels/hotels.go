package hotels

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ctrlaltreboot/igor/helper"
)

type HotelsResponse struct {
	Properties []Property `json:"properties"`
}

type Property struct {
	RoomTypes []RoomType `json:"room_types"`
}

type RoomType struct {
	Offers []Offer `json:"offers"`
}

type Offer struct {
	Charges Charges `json:"charges"`
}

type Charges struct {
	Total               Charge `json:"total"`
	PayableAtBooking    Charge `json:"payable_at_booking"`
	Discount            Charge `json:"discount"`
	BaseRate            Charge `json:"base_rate"`
	Tax                 Charge `json:"tax"`
	TaxRecovery         Charge `json:"tax_recovery"`
	ExtraOccupantCharge Charge `json:"extra_occupant_charge"`
}

type Charge struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type CheapestHandler struct {
	HotelsAPIEndpoint string
}

type cheapestRes struct {
	Properties []Property `json:"properties"`
}

func (h *CheapestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := helper.Fetch(h.HotelsAPIEndpoint)
	if err != nil {
		http.Error(w, "error reading from upstream API", http.StatusInternalServerError)
		// TODO Before going to production, we'd probably log the error here
		// to stdout/stderr and rely on logs being shipped somewhere like
		// Papertrail/Graylog/an ELK stack.
		return
	}

	var hr HotelsResponse
	if err := json.Unmarshal(b, &hr); err != nil {
		http.Error(w, "error parsing upstream API response", http.StatusInternalServerError)
		return
	}

	var res cheapestRes
	for _, p := range hr.Properties {
		res.Properties = append(res.Properties, p)
	}

	// TODO Actually figure out which are the cheapest hotels.

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "error writing JSON response", http.StatusInternalServerError)
		return
	}
}

func List(w http.ResponseWriter, r *http.Request) {
	b, err := helper.Fetch("http://127.0.0.1:5091/hotels")
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", b)
}
