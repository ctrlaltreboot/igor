package hotels

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ctrlaltreboot/igor/helper"
)

type HotelsResponse struct {
	Meta       Meta       `json:"meta"`
	Properties []Property `json:"properties"`
}

type Meta struct {
	Query `json:"query"`
}

type Query struct {
	Brand       string   `json:"brand"`
	OfferType   string   `json:"offer_type"`
	CheckIn     string   `json:"check_in"`
	CheckOut    string   `json:"check_out"`
	Adults      string   `json:"adults"`
	Children    string   `json:"children"`
	Infants     string   `json:"infants"`
	PropertyIds []string `json:"property_ids"`
}

type Property struct {
	Id        string     `json:"id"`
	RoomTypes []RoomType `json:"room_types"`
}

type RoomType struct {
	Id     string  `json:"id"`
	Offers []Offer `json:"offers"`
}

type Offer struct {
	Id                    string             `json:"id"`
	Name                  string             `json:"name"`
	Description           string             `json:"description"`
	Charges               Charges            `json:"charges"`
	CancellationPolicy    CancellationPolicy `json:"cancellation_policy"`
	Promotion             Promotion          `json:"promotion"`
	BaseIncludedOccupants int32              `json:"base_included_occupants"`
	ExtraOccupantCount    int32              `json:"extra_occupant_count"`
	Type                  string             `json:"type"`
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

type CancellationPolicy struct {
	IsRefundable bool   `json:"is_refundable"`
	Description  string `json:"description"`
}

type Promotion struct {
	PromotionCode string `json:"promotion_code"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	CampaignCode  string `json:"campaign_code"`
	Priority      int64  `json:"priority"`
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
