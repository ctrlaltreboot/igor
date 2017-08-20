package hotels

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ctrlaltreboot/igor/helper"
)

type HotelsResponse struct {
	Properties []Property `json:"properties"`
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
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Charges Charges `json:"charges"`
}

type Charges struct {
	Total Charge `json:"total"`
}

type Charge struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type propertyOffers struct {
	Properties []PropertyOffer
}

type cheapest struct {
	PropertyOffer PropertyOffer `json:"PropertyOffer"`
}

type PropertyOffer struct {
	Amount     float64 `json:"amount"`
	OfferId    string  `json:"offer_id"`
	OfferName  string  `json:"offer_name"`
	PropertyId string  `json:"property_id"`
	RoomTypeId string  `json:"room_type_id"`
}

type CheapestHandler struct {
	HotelsAPIEndpoint string
}

func toOffers(property []Property) (propertyOffers, error) {
	var pos propertyOffers
	// Get Offer slice, convert amount then save
	for _, p := range property {
		for _, rt := range p.RoomTypes {
			for _, o := range rt.Offers {
				amount, err := strconv.ParseFloat(o.Charges.Total.Amount, 64)
				if err != nil {
					return pos, err
				}

				po := PropertyOffer{
					amount,
					o.Id,
					o.Name,
					p.Id,
					rt.Id,
				}

				pos.Properties = append(pos.Properties, po)
			}
		}
	}

	return pos, nil
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

	// Get hotels JSON response, unmarshal to defined HotelsResponse struct type
	var hr HotelsResponse
	if err := json.Unmarshal(b, &hr); err != nil {
		http.Error(w, "error parsing upstream API response", http.StatusInternalServerError)
		return
	}

	// Parse the HotelsResponse in hr
	res, err := toOffers(hr.Properties)
	if err != nil {
		http.Error(w, "error reading from upstream API", http.StatusInternalServerError)
		return
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
