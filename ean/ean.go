package ean

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/ctrlaltreboot/igor/helper"
)

type EanResponse struct {
	HotelListResponse HotelListResponse `json:"HotelListResponse"`
}

type HotelListResponse struct {
	HotelList HotelList `json:"HotelList"`
}

type HotelList struct {
	HotelSummary []HotelSummary `json:"HotelSummary"`
}

type HotelSummary struct {
	HotelId             int64               `json:"hotelId"`
	RoomRateDetailsList RoomRateDetailsList `json:"RoomRateDetailsList"`
}

type RoomRateDetailsList struct {
	RoomRateDetails RoomRateDetails `json:"RoomRateDetails"`
}

type RoomRateDetails struct {
	RoomTypeCode    int64     `json:"roomTypeCode"`
	RoomDescription string    `json:"roomDescription"`
	RateInfos       RateInfos `json:"RateInfos`
}

type RateInfos struct {
	RateInfo RateInfo `json:"RateInfo"`
}

type RateInfo struct {
	ChargeableRateInfo ChargeableRateInfo `json:"ChargeableRateInfo"`
	NonRefundable      bool               `json:"nonRefundable"`
}

type ChargeableRateInfo struct {
	Total string `json:"@total"`
}

type Offer struct {
	PropertyName  string  `json:"property_name"`
	RoomTypeName  string  `json:"room_type_name"`
	OfferName     string  `json:"offer_name"`
	Amount        float64 `json:"amount"`
	NonRefundable bool    `json:"is_non_refundable"`
}

type cheapest struct {
	CheapestOffer Offer `json:"cheapest_offer"`
}

type offers struct {
	Offers []Offer
}

type CheapestHandler struct {
	EanAPIEndpoint string
}

func summaryToOffer(hs HotelSummary) Offer {
	amount, err := strconv.ParseFloat(hs.RoomRateDetailsList.RoomRateDetails.RateInfos.RateInfo.ChargeableRateInfo.Total, 64)
	if err != nil {
		panic(err)
	}

	return Offer{
		strconv.FormatInt(hs.HotelId, 10),
		strconv.FormatInt(hs.RoomRateDetailsList.RoomRateDetails.RoomTypeCode, 10),
		hs.RoomRateDetailsList.RoomRateDetails.RoomDescription,
		amount,
		hs.RoomRateDetailsList.RoomRateDetails.RateInfos.RateInfo.NonRefundable,
	}
}

func (h *CheapestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := helper.Fetch(h.EanAPIEndpoint)
	if err != nil {
		http.Error(w, "error reading from upstream API", http.StatusInternalServerError)
		// As w/ hotels, log elsewhere.
		return
	}

	var er EanResponse
	if err := json.Unmarshal(b, &er); err != nil {
		http.Error(w, "error parsing upstream API response", http.StatusInternalServerError)
		return
	}

	var o offers
	for _, hs := range er.HotelListResponse.HotelList.HotelSummary {
		o.Offers = append(o.Offers, summaryToOffer(hs))
	}

	sort.SliceStable(o.Offers, func(i, j int) bool {
		return o.Offers[i].Amount < o.Offers[j].Amount
	})

	res := cheapest{o.Offers[0]}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "error writing JSON response", http.StatusInternalServerError)
		return
	}
}

func List(w http.ResponseWriter, r *http.Request) {
	b, err := helper.Fetch("http://127.0.0.1:5092/ean")
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", b)
}
