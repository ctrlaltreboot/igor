package hotels

import (
	"encoding/json"
	"sort"
	"strconv"

	"igor/helper"
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
	Id                 string             `json:"id"`
	Name               string             `json:"name"`
	Charges            Charges            `json:"charges"`
	CancellationPolicy CancellationPolicy `json:"cancellation_policy"`
}

type Charges struct {
	Total Charge `json:"total"`
}

type Charge struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type CancellationPolicy struct {
	IsNonRefundable bool `json:"is_nonrefundable"`
}

type propertyOffers struct {
	Offers []PropertyOffer
}

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
					p.Id,
					rt.Id,
					o.Name,
					amount,
					o.CancellationPolicy.IsNonRefundable,
				}

				pos.Offers = append(pos.Offers, po)
			}
		}
	}

	return pos, nil
}

func GetCheapest(url string) PropertyOffer {
	b, err := helper.Fetch(url)

	// Get hotels JSON response, unmarshal to defined HotelsResponse struct type
	var hr HotelsResponse
	if err := json.Unmarshal(b, &hr); err != nil {
		panic(err)
	}

	// Parse the HotelsResponse in hr
	pos, err := toOffers(hr.Properties)
	if err != nil {
		panic(err)
	}

	sort.SliceStable(pos.Offers, func(i, j int) bool {
		return pos.Offers[i].Amount < pos.Offers[j].Amount
	})

	return pos.Offers[0]
}
